package main

import (
	"dgl-kafka-publisher/pkg/handler"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

func startConsumer(servers []string) {
	topicName := "testkafka"

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

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			slog.Info("Received message", msg.Value)
		}
	}

}
func main() {
	servers := []string{"kafka:9092", "localhost:9092"}
	// servers := []string{"192.168.1.143:9092", "kafka:9092", "localhost:9092"}
	saramaClient, err := sarama.NewClient(servers, nil)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	app := fiber.New()
	go startConsumer(servers)
	version := app.Group("/v1")
	api := version.Group("/api")

	topicHandler := handler.NewTopicHandler(saramaClient)
	intuitionHandler := handler.NewMockupIntutionHandler()
	api.Get("/topics", topicHandler.GetTopics)
	api.Post("/intuition", intuitionHandler.PostToInstution)
	app.Listen(":3000")

	// servers := []string{"kafka:9092"}
	// topicName := "testkafka"
	// consumer, err := sarama.NewConsumer(servers, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// defer consumer.Close()

	// partitionConsumer, err := consumer.ConsumePartition(topicName, 0, sarama.OffsetNewest)
	// if err != nil {
	// 	panic(err)
	// }

	// defer partitionConsumer.Close()
	// for {
	// 	select {
	// 	case msg := <-partitionConsumer.Messages():
	// 		println("Received message", string(msg.Value))
	// 	}
	// }

}
