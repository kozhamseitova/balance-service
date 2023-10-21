package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) InitRoutes(router *fiber.App) {
	api := router.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("test")
	})

}