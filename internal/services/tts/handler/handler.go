package handler

import (
	"net/http"

	"github.com/av-ugolkov/lingua-ai/internal/services/tts"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	svc *tts.Service
}

func Create(r *fiber.App, s *tts.Service) {
	h := Handler{
		svc: s,
	}

	r.Get("/tts", h.GetAudio)
}

func (h *Handler) GetAudio(c *fiber.Ctx) error {
	text := c.Query("text")
	lang := c.Query("lang")
	_ = h.svc.GetAudio(text, lang)
	return c.SendStatus(http.StatusOK)
}
