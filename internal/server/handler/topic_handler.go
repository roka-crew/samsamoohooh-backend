package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
)

type TopicHandler struct {
}

func NewTopicHandler() *TopicHandler {
	handler := &TopicHandler{}
	return handler
}

func (h TopicHandler) CreateTopic(c *fiber.Ctx) error {
	var (
		_ domain.CreateTopicRequest
		_ domain.CreateTopicResponse
		_ error
	)
	return nil
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
