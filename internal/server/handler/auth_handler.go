package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/server"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/middleware"
	"github.com/roka-crew/samsamoohooh-backend/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(
	server *server.Server,
	authService *service.AuthService,
	authMiddleware *middleware.AuthMiddleware,
) *AuthHandler {
	handler := &AuthHandler{
		authService: authService,
	}

	auth := server.Group("/auth")
	{
		auth.Post("/login", handler.Login)
		auth.Post("/validate", authMiddleware.Authenticate, handler.Validate)
	}

	return handler
}

func (h AuthHandler) Login(c *fiber.Ctx) error {
	var (
		request  domain.LoginRequest
		response domain.LoginResponse
		err      error
	)

	response, err = h.authService.Login(c.Context(), request)

	switch {
	case err == nil:
		return c.Status(fiber.StatusOK).JSON(response)
	default:
		return err
	}
}

func (h AuthHandler) Validate(c *fiber.Ctx) error {
	var (
		request  domain.ValidateRequest
		response domain.ValidateResponse
		err      error
	)

	if err = c.ReqHeaderParser(&request); err != nil {
		return err
	}

	response, err = h.authService.Validate(c.Context(), request)

	switch {
	case err == nil:
		return c.Status(fiber.StatusOK).JSON(response)
	default:
		return err
	}
}
