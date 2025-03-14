package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/handler/validator"
	"github.com/roka-crew/samsamoohooh-backend/internal/service"
	"github.com/roka-crew/samsamoohooh-backend/pkg/apperr"
	"github.com/roka-crew/samsamoohooh-backend/pkg/server/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(
	server *http.Server,
	userService *service.UserService,
) *UserHandler {
	handler := &UserHandler{
		userService: userService,
	}

	users := server.Group("/users")
	{
		users.Post("/", handler.CreateUser)
	}

	return handler
}

// CreateUser godoc
// @Tags users
// @Summary 새로운 사용자 생성 ✅
// @Description 새로운 사용자 생성 API
// @Accept json
// @Produce json
// @Param CreateUserRequest body domain.CreateUserRequest true "새로운 사용자 정보"
// @Success 201 {object} domain.CreateUserResponse "성공적으로 사용자를 생성한 경우"
// @Failure 409 {object} apperr.Apperr "이미 존재하는 사용자가 있는 경우"
// @Router /users [post]
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
