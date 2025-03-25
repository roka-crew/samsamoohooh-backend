package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/validator"
	"github.com/roka-crew/samsamoohooh-backend/internal/service"
)

type TopicHandler struct {
	topicService *service.TopicService
}

func NewTopicHandler(
	topicService *service.TopicService,
) *TopicHandler {
	return &TopicHandler{
		topicService: topicService,
	}
}

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

func (h TopicHandler) ListTopics(c *fiber.Ctx) error {
	var (
		_ domain.ListTopicsRequest
		_ domain.ListTopicsResponse
		_ error
	)
	return nil
}

func (h TopicHandler) PatchTopic(c *fiber.Ctx) error {
	var (
		_ domain.PatchTopicRequest
		_ error
	)
	return nil
}

func (h TopicHandler) DeleteTopic(c *fiber.Ctx) error {
	var (
		_ domain.DeleteTopicRequest
		_ error
	)
	return nil
}
