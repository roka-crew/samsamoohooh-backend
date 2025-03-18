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
