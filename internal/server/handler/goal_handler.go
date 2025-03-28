package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/server"
)

type GoalHandler struct {
}

func NewGoalHander(
	server *server.Server,
) *GoalHandler {
	handler := &GoalHandler{}

	goals := server.Group("/goals", server.AuthMiddleware.Authenticate)
	{
		goals.Post("/", handler.CreateGoal)
		goals.Get("/", handler.ListGoals)
		goals.Patch("/", handler.PatchGoal)
		goals.Delete("/", handler.DeleteGoal)
	}

	return handler
}

// CreateGoal godoc
// @Tags goals
// @Summary 새로운 목표 생성 ✅
// @Accept json
// @Produce json
// @Param CreateGoalReqeust body domain.CreateGoalRequest true "생성할 목표 정보"
// @Success 201 {object} domain.CreateGoalResponse "성공적으로 목표를 생성한 경우"
// @Router /goals [post]
// @Security BearerAuth
func (h GoalHandler) CreateGoal(c *fiber.Ctx) error {
	var (
		_ domain.CreateGoalRequest
		_ domain.CreateGoalResponse
		_ error
	)

	return nil
}

// ListGoals godoc
// @Tags goals
// @Summary 목표 목록 조회 ✅
// @Accept json
// @Produce json
// @Param ListGoalsRequest query domain.ListGoalsRequest true "조회할 목표 정보"
// @Success 200 {object} domain.ListGoalsResponse "목표 목록 조회 성공"
// @Router /goals [get]
// @Security BearerAuth
func (h GoalHandler) ListGoals(c *fiber.Ctx) error {
	var (
		_ domain.ListGoalsRequest
		_ domain.ListGoalsResponse
		_ error
	)

	return nil
}

// PatchGoal godoc
// @Tags goals
// @Summary 목표 정보 수정 ✅
// @Accept json
// @Produce json
// @Param PatchGoalRequest body domain.PatchGoalRequest true "수정할 목표 정보"
// @Success 204
// @Router /goals [patch]
// @Security BearerAuth
func (h GoalHandler) PatchGoal(c *fiber.Ctx) error {
	var (
		_ domain.PatchGoalRequest
		_ error
	)

	return nil
}

// DeleteGoal godoc
// @Tags goals
// @Summary 목표 삭제 ✅
// @Accept json
// @Produce json
// @Param DeleteGoalRequest body domain.DeleteGoalRequest true "삭제할 목표 정보"
// @Success 204
// @Router /goals [delete]
// @Security BearerAuth
func (h GoalHandler) DeleteGoal(c *fiber.Ctx) error {
	var (
		_ domain.DeleteGoalRequest
		_ error
	)

	return nil
}
