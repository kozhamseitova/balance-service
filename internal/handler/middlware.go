package handler

import (
	"github.com/gofiber/fiber/v2"
)

const (
	traceID = "traceID"
)

func (h *Handler) generateTraceId(c *fiber.Ctx) error {
	traceId := c.Get(traceID)
	if traceId == "" {
		context := h.logger.SetTraceID(c.UserContext())
		c.SetUserContext(context)
	}

	return c.Next()
}