package server

import (
	"context"

	"github.com/kozhamseitova/balance-service/internal/config"
	"github.com/kozhamseitova/balance-service/internal/handler"
	"github.com/kozhamseitova/balance-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func New(cfg *config.Config, logger logger.Logger, handler *handler.Handler) *Server {
	app := fiber.New()

	handler.InitRoutes(app)

	return &Server{
		app: app,
	}
}

func (s *Server) Run(port string) error {
	return s.app.Listen(":" + port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}