package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"

	_ "github.com/roka-crew/samsamoohooh-backend/api"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/middleware"
	"github.com/roka-crew/samsamoohooh-backend/pkg/apperr"
	"github.com/roka-crew/samsamoohooh-backend/pkg/config"
	"go.uber.org/fx"
)

type Server struct {
	*fiber.App
	AuthMiddleware *middleware.AuthMiddleware
}

func NewServer(
	config *config.Config,
	lifeCycle fx.Lifecycle,

	AuthMiddleware *middleware.AuthMiddleware,
) *Server {
	server := &Server{
		App: fiber.New(fiber.Config{
			AppName:      config.Name,
			ErrorHandler: errorHandler,
		}),

		AuthMiddleware: AuthMiddleware,
	}

	lifeCycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.Listen(config.Listen); err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown()
		},
	})

	server.Get("/swagger/*", swagger.HandlerDefault)
	server.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	server.Use(recover.New())

	return server
}

var errorHandler = func(c *fiber.Ctx, err error) error {
	var appError *apperr.Apperr
	if errors.As(err, &appError) && appError != nil {
		return c.Status(appError.Status).JSON(appError)
	}

	var internalError *apperr.InternalError
	if errors.As(err, &internalError) && internalError != nil {
		fmt.Println("internal error: ", err)
		fmt.Println("StackTrace()")
		fmt.Println(internalError.StackTrace(func(file, _ string, line int) string {
			return fmt.Sprintf("\t%s:%d", file, line)
		}))

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return fiber.DefaultErrorHandler(c, err)
}
