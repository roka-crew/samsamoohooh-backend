package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/server"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/ctxutil"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/validator"
	"github.com/roka-crew/samsamoohooh-backend/internal/service"
)

type GoalHandler struct {
	goalService *service.GoalService
}

func NewGoalHandler(
	goalService *service.GoalService,
	server *server.Server,
) *GoalHandler {
	handler := &GoalHandler{
		goalService: goalService,
	}

	goals := server.Group("/goals", server.AuthMiddleware.Authenticate)
	{
		goals.Post("/", handler.CreateGoal)
		goals.Get("/", handler.ListGoals)
		goals.Patch("/:goalID", handler.PatchGoal)
		goals.Delete("/", handler.DeleteGoal)
	}

	return handler
}

// CreateGoal godoc
//
//	@Tags		goals
//	@Summary	새로운 목표 생성 ✅
//	@Accept		json
//	@Produce	json
//	@Param		CreateGoalRequest	body		domain.CreateGoalRequest	true	"생성할 목표 정보"
//	@Success	201					{object}	domain.CreateGoalResponse	"성공적으로 목표를 생성한 경우"
//	@Router		/goals [post]
//	@Security	BearerAuth
func (h GoalHandler) CreateGoal(c *fiber.Ctx) error {
	var (
		request  domain.CreateGoalRequest
		response domain.CreateGoalResponse
		err      error
	)

	if err = c.BodyParser(&request); err != nil {
		return err
	}

	if request.RequestUserID, err = ctxutil.GetUserID(c); err != nil {
		return err
	}

	if err = validator.Validate(&request); err != nil {
		return err
	}

	response, err = h.goalService.CreateGoal(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// ListGoals godoc
//
//	@Tags		goals
//	@Summary	목표 목록 조회 ✅
//	@Accept		json
//	@Produce	json
//	@Param		ListGoalsRequest	query		domain.ListGoalsRequest		true	"조회할 목표 정보"
//	@Success	200					{object}	domain.ListGoalsResponse	"목표 목록 조회 성공"
//	@Router		/goals [get]
//	@Security	BearerAuth
func (h GoalHandler) ListGoals(c *fiber.Ctx) error {
	var (
		request  domain.ListGoalsRequest
		response domain.ListGoalsResponse
		err      error
	)

	if err = c.QueryParser(&request); err != nil {
		return err
	}

	if request.RequestUserID, err = ctxutil.GetUserID(c); err != nil {
		return err
	}

	if err = validator.Validate(&request); err != nil {
		return err
	}

	response, err = h.goalService.ListGoals(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// PatchGoal godoc
//
//	@Tags		goals
//	@Summary	목표 정보 수정 ✅
//	@Accept		json
//	@Produce	json
//	@Param		goalID				path	string					true	"수정할 목표 ID"
//	@Param		PatchGoalRequest	body	domain.PatchGoalRequest	true	"수정할 목표 정보"
//	@Success	204
//	@Router		/goals/{goalID} [patch]
//	@Security	BearerAuth
func (h GoalHandler) PatchGoal(c *fiber.Ctx) error {
	var (
		request domain.PatchGoalRequest
		err     error
	)

	if err = c.ParamsParser(&request); err != nil {
		return err
	}

	if err = c.BodyParser(&request); err != nil {
		return err
	}

	if request.RequestUserID, err = ctxutil.GetUserID(c); err != nil {
		return err
	}

	if err = validator.Validate(&request); err != nil {
		return err
	}

	err = h.goalService.PatchGoal(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteGoal godoc
//
//	@Tags		goals
//	@Summary	목표 삭제 ✅
//	@Accept		json
//	@Produce	json
//	@Param		DeleteGoalRequest	body	domain.DeleteGoalRequest	true	"삭제할 목표 정보"
//	@Success	204
//	@Router		/goals [delete]
//	@Security	BearerAuth
func (h GoalHandler) DeleteGoal(c *fiber.Ctx) error {
	var (
		request domain.DeleteGoalRequest
		err     error
	)

	if err = c.BodyParser(&request); err != nil {
		return err
	}

	if request.RequestUserID, err = ctxutil.GetUserID(c); err != nil {
		return err
	}

	if err = validator.Validate(&request); err != nil {
		return err
	}

	err = h.goalService.DeleteGoal(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
