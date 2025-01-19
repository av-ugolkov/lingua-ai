package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

func Create(r *fiber.App) {
	h := Handler{}

	r.Get("/health", h.GetHealth)
}

func (h *Handler) GetHealth(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}
