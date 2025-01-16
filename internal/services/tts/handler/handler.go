package handler

import (
	"net/http"

	"github.com/av-ugolkov/lingua-ai/internal/services/tts"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	ctx := c.Context()

	var query struct {
		ID   uuid.UUID `json:"id"`
		Text string    `json:"text"`
		Lang string    `json:"lang"`
	}

	err := c.QueryParser(&query)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	data, err := h.svc.GetAudio(ctx, query.ID, query.Text, query.Lang)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	c.Context().SetContentType("audio/wav")
	return c.Status(http.StatusOK).Send(data)
}
