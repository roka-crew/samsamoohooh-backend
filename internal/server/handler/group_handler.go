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

type GroupHandler struct {
	groupService *service.GroupService
}

func NewGroupHandler(
	server *server.Server,
	authMiddleware *middleware.AuthMiddleware,
	groupService *service.GroupService,
) *GroupHandler {
	handler := &GroupHandler{
		groupService: groupService,
	}

	groups := server.Group("/groups", authMiddleware.Authenticate)
	{
		groups.Post("/", handler.CreateGroup)
		groups.Get("/", handler.ListGroups)
		groups.Patch("/:group-id", handler.PatchGroup)

		groups.Post("/join", handler.JoinGroup)
		groups.Post("/leave", handler.LeaveGroup)
	}

	return handler
}

func (h GroupHandler) CreateGroup(c *fiber.Ctx) error {
	var (
		request  domain.CreateGroupRequest
		response domain.CreateGroupResponse
		err      error
	)

	if err = c.BodyParser(&request); err != nil {
		return err
	}

	if err = validator.Validate(request); err != nil {
		return err
	}

	request.UserID, err = ctxutil.GetUserID(c)
	if err != nil {
		return err
	}

	response, err = h.groupService.CreateGroup(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h GroupHandler) ListGroups(c *fiber.Ctx) error {
	var (
		request  domain.ListGroupsRequest
		response domain.ListGroupsResponse
		err      error
	)

	if err = c.QueryParser(&request); err != nil {
		return err
	}

	if err = validator.Validate(request); err != nil {
		return err
	}

	response, err = h.groupService.ListGroups(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h GroupHandler) PatchGroup(c *fiber.Ctx) error {
	var (
		request domain.PatchGroupRequest
		err     error
	)

	if err = c.ParamsParser(&request); err != nil {
		return err
	}

	if err = c.BodyParser(&request); err != nil {
		return err
	}

	if err = validator.Validate(&request); err != nil {
		return err
	}

	err = h.groupService.PatchGroup(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h GroupHandler) JoinGroup(c *fiber.Ctx) error {
	var (
		request domain.JoinGroupRequest
		err     error
	)

	if err = c.BodyParser(&request); err != nil {
		return err
	}

	if err = validator.Validate(&request); err != nil {
		return err
	}

	if request.UserID, err = ctxutil.GetUserID(c); err != nil {
		return err
	}

	err = h.groupService.JoinGroup(c.Context(), domain.JoinGroupRequest{})
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h GroupHandler) LeaveGroup(c *fiber.Ctx) error {
	var (
		request domain.LeaveGroupRequest
		err     error
	)

	if err = c.BodyParser(&request); err != nil {
		return err
	}

	if err = validator.Validate(&request); err != nil {
		return err
	}

	if request.UserID, err = ctxutil.GetUserID(c); err != nil {
		return err
	}

	err = h.groupService.LeaveGroup(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
