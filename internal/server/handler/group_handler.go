package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/server"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/ctxutil"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/validator"
	"github.com/roka-crew/samsamoohooh-backend/internal/service"
)

type GroupHandler struct {
	groupService *service.GroupService
}

func NewGroupHandler(
	server *server.Server,
	groupService *service.GroupService,
) *GroupHandler {
	handler := &GroupHandler{
		groupService: groupService,
	}

	groups := server.Group("/groups", server.AuthMiddleware.Authenticate)
	{
		groups.Post("/", handler.CreateGroup)
		groups.Get("/", handler.ListGroups)
		groups.Patch("/:groupID", handler.PatchGroup)

		groups.Post("/join", handler.JoinGroup)
		groups.Post("/leave", handler.LeaveGroup)
	}

	return handler
}

// CreateGroup godoc
//
//	@Tags		groups
//	@Summary	새로운 모임 생성 ✅
//	@Accept		json
//	@Produce	json
//	@Param		CreateGroupRequest	body		domain.CreateGroupRequest	true	"생성할 모임 정보"
//	@Success	201					{object}	domain.CreateGroupResponse	"성공적으로 모임을 생성한 경우"
//	@Router		/groups [post]
//	@Security	BearerAuth
func (h GroupHandler) CreateGroup(c *fiber.Ctx) error {
	var (
		request  domain.CreateGroupRequest
		response domain.CreateGroupResponse
		err      error
	)

	if err = c.BodyParser(&request); err != nil {
		return err
	}

	request.RequestUserID, err = ctxutil.GetUserID(c)
	if err != nil {
		return err
	}

	if err = validator.Validate(request); err != nil {
		return err
	}

	response, err = h.groupService.CreateGroup(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// ListGroups godoc
//
//	@Tags		groups
//	@Summary	모임 리스트 ✅
//	@Accept		json
//	@Produce	json
//	@Param		limit	query		int							false	"조회할 모임 개수"
//	@Success	200		{object}	domain.ListGroupsResponse	"성공적으로 모임 리스트를 조회한 경우"
//	@Router		/groups [get]
//	@Security	BearerAuth
func (h GroupHandler) ListGroups(c *fiber.Ctx) error {
	var (
		request  domain.ListGroupsRequest
		response domain.ListGroupsResponse
		err      error
	)

	if err = c.QueryParser(&request); err != nil {
		return err
	}

	if request.RequesterID, err = ctxutil.GetUserID(c); err != nil {
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

// PatchGroup godoc
//
//	@Tags		groups
//	@Summary	모임 정보 수정 ✅
//	@Accept		json
//	@Produce	json
//	@Param		group-id			path	string						true	"수정할 모임 ID"
//	@Param		PatchGroupRequest	body	domain.PatchGroupRequest	true	"수정할 모임 정보"
//	@Success	204
//	@Router		/groups/{group-id} [patch]
//	@Security	BearerAuth
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

	if request.RequestUserID, err = ctxutil.GetUserID(c); err != nil {
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

// JoinGroup godoc
//
//	@Tags		groups
//	@Summary	모임 참가 ✅
//	@Accept		json
//	@Produce	json
//	@Param		JoinGroupRequest	body	domain.JoinGroupRequest	true	"참여할 모임 정보"
//	@Success	204
//	@Router		/groups/join [post]
//	@Security	BearerAuth
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

	if request.RequesterID, err = ctxutil.GetUserID(c); err != nil {
		return err
	}

	err = h.groupService.JoinGroup(c.Context(), domain.JoinGroupRequest{})
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// LeaveGroup godoc
//
//	@Tags		groups
//	@Summary	모임 탈퇴 ✅
//	@Accept		json
//	@Produce	json
//	@Param		LeaveGroupRequest	body	domain.LeaveGroupRequest	true	"탈퇴할 모임 정보"
//	@Success	204
//	@Router		/groups/leave [post]
//	@Security	BearerAuth
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

	if request.RequesterID, err = ctxutil.GetUserID(c); err != nil {
		return err
	}

	err = h.groupService.LeaveGroup(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
