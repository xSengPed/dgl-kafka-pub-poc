package handler

import (
	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
)

type TopicHandler struct {
	saramaClient sarama.Client
}

func NewTopicHandler(saramaClient sarama.Client) *TopicHandler {
	return &TopicHandler{saramaClient: saramaClient}
}

func (h *TopicHandler) GetTopics(ctx *fiber.Ctx) error {
	topics, err := h.saramaClient.Topics()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"topics": topics})
}
