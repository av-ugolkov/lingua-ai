package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/av-ugolkov/lingua-ai/internal/closer"
	"github.com/av-ugolkov/lingua-ai/internal/config"
	ttsService "github.com/av-ugolkov/lingua-ai/internal/services/tts"
	ttsHandler "github.com/av-ugolkov/lingua-ai/internal/services/tts/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jsoniter "github.com/json-iterator/go"
)

func ServerStart(cfg *config.Config) {
	router := fiber.New(fiber.Config{
		AppName:      "Lingua AI",
		JSONEncoder:  jsoniter.Marshal,
		JSONDecoder:  jsoniter.Unmarshal,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			slog.Error(err.Error())

			if e, ok := err.(*fiber.Error); ok {
				return c.Status(e.Code).JSON(fiber.Map{
					"error": e.Message,
				})
			}

			// Для других ошибок — возвращаем 500
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		},
	})
	router.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(cfg.Service.AllowedOrigins, ","),
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowCredentials: true,
		AllowHeaders:     "Authorization,Content-Type,Fingerprint",
	}))

	initServer(cfg, router)

	address := fmt.Sprintf(":%d", cfg.Service.Port)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	slog.Info("start server")
	go func() {
		err := router.Listen(address)

		if err != nil {
			switch {
			case errors.Is(err, http.ErrServerClosed):
				slog.Warn("server shutdown")
			default:
				slog.Error(fmt.Sprintf("server returned an err: %v\n", err.Error()))
				return
			}
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := router.ShutdownWithContext(ctx); err != nil {
		slog.Error(fmt.Sprintf("server shutdown returned an err: %v\n", err))
	}

	err := closer.Close(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("closer: %v\n", err.Error()))
	}

	slog.Info("final")
}

func initServer(cfg *config.Config, r *fiber.App) {
	slog.Info("create services")
	ttsSvc := ttsService.New(cfg.Tts)
	closer.Add(ttsSvc.Close)

	slog.Info("create handlers")
	ttsHandler.Create(r, ttsSvc)

	slog.Info("end init services")
}
