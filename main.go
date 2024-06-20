package main

import (
	"github.com/IBM/sarama"
)

func main() {

	servers := []string{"kafka:9092"}
	topicName := "testkafka"
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
			println("Received message", string(msg.Value))
		}
	}

}
