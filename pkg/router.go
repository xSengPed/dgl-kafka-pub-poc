package router

import (
	"dgl-kafka-publisher/pkg/handler"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

func startConsumer(servers []string) {
	topicName := "dgl_kafka_test"

	slog.Info("Starting consumer")
	consumer, err := sarama.NewConsumer(servers, nil)
	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topicName, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		slog.Info("Received message", msg.Value)
	}
}
func NewRouter(saramaClient sarama.Client, servers []string) *fiber.App {
	app := fiber.New()
	go startConsumer(servers)
	version := app.Group("/v1")
	api := version.Group("/api")

	topicHandler := handler.NewTopicHandler(saramaClient)
	intuitionHandler := handler.NewMockupIntutionHandler()
	api.Get("/topics", topicHandler.GetTopics)
	api.Post("/intuition", intuitionHandler.PostToInstution)
	api.Post("/intuition-v2", intuitionHandler.PostToInstutionV2)
	api.Get("/gen-bcl", intuitionHandler.GenerateBCLMock)
	return app
}
