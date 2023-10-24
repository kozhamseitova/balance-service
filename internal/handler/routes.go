package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swaggo/fiber-swagger"

	_ "github.com/kozhamseitova/balance-service/docs"
)

func (h *Handler) InitRoutes(router *fiber.App) {

	router.Get("/swagger/*", fiberSwagger.WrapHandler)

	api := router.Group("/api")
	api.Use(h.generateTraceId)

	v1 := api.Group("/v1")
	v1.Get("/report", h.getReport)

	balance := v1.Group("/balance")

	balance.Get(":id", h.getBalance)
	balance.Post("/credit", h.deposit)
	balance.Post("/reserve", h.reserveFunds)
	balance.Post("/revenue" ,h.recognizeRevenue)
	balance.Delete("/reserve", h.cancelReservation)

}