package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/server"
	"github.com/roka-crew/samsamoohooh-backend/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(
	server *server.Server,
	authService *service.AuthService,
) *AuthHandler {
	handler := &AuthHandler{
		authService: authService,
	}

	auth := server.Group("/auth")
	{
		auth.Post("/login", handler.Login)
		auth.Post("/validate", handler.Validate)
	}

	return handler
}

// Login godoc
//
//	@Tags		auth
//	@Summary	로그인 ✅
//	@Accept		json
//	@Produce	json
//	@Param		LoginRequest	body		domain.LoginRequest		true	"로그인에 필요한 정보"
//	@Success	200				{object}	domain.LoginResponse	"성공적으로 로그인을 성공한 경우"
//	@Router		/auth/login [post]
func (h AuthHandler) Login(c *fiber.Ctx) error {
	var (
		request  domain.LoginRequest
		response domain.LoginResponse
		err      error
	)

	if err = c.BodyParser(&request); err != nil {
		return err
	}

	response, err = h.authService.Login(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Validate godoc
//
//	@Tags		auth
//	@Summary	유효성 검증 ✅
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer 토큰"
//	@Success	200				{object}	domain.ValidateResponse	"성공적으로 유효성 검증을 성공한 경우"
//	@Router		/auth/validate [post]
//	@Security	Authorization
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
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
