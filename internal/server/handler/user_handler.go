package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/server"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/ctxutil"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/middleware"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/validator"
	"github.com/roka-crew/samsamoohooh-backend/internal/service"
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
		users.Post("/random", handler.CreateRanomUser)
		users.Patch("/", authMiddleware.Authenticate, handler.PatchUser)
		users.Delete("/", authMiddleware.Authenticate, handler.DeleteUser)
	}

	return handler
}

// CreateUser godoc
//
//	@Tags		users
//	@Summary	새로운 사용자 생성 ✅
//	@Accept		json
//	@Produce	json
//	@Param		CreateUserRequest	body		domain.CreateUserRequest	true	"생성할 사용자 정보"
//	@Success	201					{object}	domain.CreateUserResponse	"성공적으로 사용자를 생성한 경우"
//	@Router		/users [post]
//	@Security	BearerAuth
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
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// CreateRandomUser
//
//	@Tags		users
//	@Summary	새로운 임의 사용자 생성 ✅
//	@Accept		json
//	@Produce	json
//	@Success	201	{object}	domain.CreateRandomUserResponse	"성공적으로 임의의 사용자를 생성한 경우"
//	@Router		/users/random [post]
//	@Security	BearerAuth
func (h UserHandler) CreateRanomUser(c *fiber.Ctx) error {
	var (
		response domain.CreateRandomUserResponse
		err      error
	)

	response, err = h.userService.CreateRandomUser(c.Context())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// PatchUser godoc
//
//	@Tags		users
//	@Summary	사용자 정보 수정 ✅
//	@Accept		json
//	@Produce	json
//	@Param		PatchUserRequest	body	domain.PatchUserRequest	true	"수정할 사용자 정보"
//	@Success	204
//	@Router		/users [patch]
//	@Security	BearerAuth
func (h UserHandler) PatchUser(c *fiber.Ctx) error {
	var (
		request domain.PatchUserRequest
		err     error
	)

	if err = c.BodyParser(&request); err != nil {
		return err
	}

	request.RequestUserID, err = ctxutil.GetUserID(c)
	if err != nil {
		return err
	}

	if err = validator.Validate(&request); err != nil {
		return err
	}

	err = h.userService.PatchUser(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteUser godoc
//
//	@Tags		users
//	@Summary	사용자 삭제 ✅
//	@Router		/users [delete]
//	@Success	204
//	@Security	BearerAuth
func (h UserHandler) DeleteUser(c *fiber.Ctx) error {
	var (
		request domain.DeleteUserRequest
		err     error
	)

	request.RequestUserID, err = ctxutil.GetUserID(c)
	if err != nil {
		return err
	}

	if err = validator.Validate(&request); err != nil {
		return err
	}

	err = h.userService.DeleteUser(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
