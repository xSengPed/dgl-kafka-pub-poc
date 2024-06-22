package main

import (
	router "dgl-kafka-publisher/pkg"

	"github.com/IBM/sarama"
	"golang.org/x/exp/slog"
)

func main() {
	servers := []string{"127.0.0.1:29092"}

	saramaClient, err := sarama.NewClient(servers, nil)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	app := router.NewRouter(saramaClient, servers)

	app.Listen(":3000")

}
