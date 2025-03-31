package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/server"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/ctxutil"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/validator"
	"github.com/roka-crew/samsamoohooh-backend/internal/service"
)

type TopicHandler struct {
	topicService *service.TopicService
}

func NewTopicHandler(
	server *server.Server,
	topicService *service.TopicService,
) *TopicHandler {
	handler := &TopicHandler{
		topicService: topicService,
	}

	topics := server.Group("/topics", server.AuthMiddleware.Authenticate)
	{
		topics.Post("/", handler.CreateTopic)
		topics.Get("/", handler.ListTopics)
		topics.Patch("/:topicID", handler.PatchTopic)
		topics.Delete("/", handler.DeleteTopic)
	}

	return handler
}

// CreateTopic godoc
//
//	@Tags		topics
//	@Summary	새로운 주제 생성 ✅
//	@Accept		json
//	@Produce	json
//	@Param		CreateTopicReqeust	body		domain.CreateTopicRequest	true	"생성할 주제 정보"
//	@Success	201					{object}	domain.CreateTopicResponse	"성공적으로 주제를 생성한 경우"
//	@Router		/topics [post]
//	@Security	BearerAuth
func (h TopicHandler) CreateTopic(c *fiber.Ctx) error {
	var (
		request  domain.CreateTopicRequest
		response domain.CreateTopicResponse
		err      error
	)

	if err = c.BodyParser(&request); err != nil {
		return err
	}

	if err = validator.Validate(&request); err != nil {
		return err
	}

	response, err = h.topicService.CreateTopic(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// ListTopics godoc
//
//	@Tags		topics
//	@Summary	주제 목록 조회 ✅
//	@Accept		json
//	@Produce	json
//	@Param		ListTopicsRequest	query		domain.ListTopicsRequest	true	"조회할 주제 정보"
//	@Success	200					{object}	domain.ListTopicsResponse	"주제 목록 조회 성공"
//	@Router		/topics [get]
//	@Security	BearerAuth
func (h TopicHandler) ListTopics(c *fiber.Ctx) error {
	var (
		request  domain.ListTopicsRequest
		response domain.ListTopicsResponse
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

	response, err = h.topicService.ListTopics(c.Context(), request)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// PatchTopic godoc
//
//	@Tags		topics
//	@Summary	주제 정보 수정 ✅
//	@Accept		json
//	@Produce	json
//	@Param		goalID				path	string						true	"수정할 목표 ID"
//	@Param		PatchTopicRequest	body	domain.PatchTopicRequest	true	"수정할 주제 정보"
//	@Success	204
//	@Router		/topics/{topicID} [patch]
//	@Security	BearerAuth
func (h TopicHandler) PatchTopic(c *fiber.Ctx) error {
	var (
		request domain.PatchTopicRequest
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

	err = h.topicService.PatchTopic(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteTopic godoc
//
//	@Tags		topics
//	@Summary	주제 삭제 ✅
//	@Accept		json
//	@Produce	json
//	@Param		DeleteTopicRequest	body	domain.DeleteTopicRequest	true	"삭제할 주제 정보"
//	@Success	204
//	@Router		/topics [delete]
//	@Security	BearerAuth
func (h TopicHandler) DeleteTopic(c *fiber.Ctx) error {
	var (
		request domain.DeleteTopicRequest
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

	err = h.topicService.DeleteTopic(c.Context(), request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
