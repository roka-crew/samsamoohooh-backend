package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/server"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/ctxutil"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/middleware"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/validator"
	"github.com/roka-crew/samsamoohooh-backend/internal/service"
	"github.com/roka-crew/samsamoohooh-backend/pkg/apperr"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(
	server *server.Server,
	userService *service.UserService,
	authMiddleware *middleware.AuthMiddleware,
) *UserHandler {
	handler := &UserHandler{
		userService: userService,
	}

	users := server.Group("/users")
	{
		users.Post("/", handler.CreateUser)
		users.Patch("/", authMiddleware.Authenticate, handler.PatchUser)
		users.Delete("/", authMiddleware.Authenticate, handler.DeleteUser)
	}

	return handler
}

// CreateUser godoc
//
//	@Tags			users
//	@Summary		새로운 사용자 생성 ✅
//	@Description	|message|status|description|
//	@Description	|---|---|---|
//	@Description	|`ERR_UESR_DUPLICATE`|409|이미 존재하는 사용자가 있는 경우|
//	@Accept			json
//	@Produce		json
//	@Param			CreateUserRequest	body		domain.CreateUserRequest	true	"새로운 사용자 정보"
//	@Success		201					{object}	domain.CreateUserResponse	"성공적으로 사용자를 생성한 경우"
//	@Router			/users [post]
func (h UserHandler) CreateUser(c *fiber.Ctx) error {
	var (
		request  domain.CreateUserRequest
		response domain.CreateUserResponse
		err      error
	)

	if err = c.BodyParser(&request); err != nil {
		return err
	}

	if err = validator.Validate(&request); err != nil {
		return err
	}

	response, err = h.userService.CreateUser(c.Context(), request)

	switch {
	case err == nil:
		return c.Status(fiber.StatusCreated).JSON(response)
	case errors.Is(err, domain.ErrUserAlreadyExists):
		return err.(*apperr.Apperr).WithStatus(fiber.StatusConflict)
	default:
		return err
	}
}

func (h UserHandler) PatchUser(c *fiber.Ctx) error {
	var (
		request domain.PatchUserRequest
		err     error
	)

	if err = c.BodyParser(&request); err != nil {
		return err
	}

	request.UserID, err = ctxutil.GetUserID(c)
	if err != nil {
		return err
	}

	err = h.userService.PatchUser(c.Context(), request)

	switch {
	case err == nil:
		return c.SendStatus(fiber.StatusNoContent)
	case errors.Is(err, domain.ErrUserNotFound):
		return err.(*apperr.Apperr).WithStatus(fiber.StatusNotFound)
	default:
		return err
	}
}

func (h UserHandler) DeleteUser(c *fiber.Ctx) error {
	var (
		request domain.DeleteUserRequest
		err     error
	)

	request.UserID, err = ctxutil.GetUserID(c)
	if err != nil {
		return err
	}

	err = h.userService.DeleteUser(c.Context(), request)

	switch {
	case err == nil:
		return c.SendStatus(fiber.StatusNoContent)
	default:
		return err
	}
}
