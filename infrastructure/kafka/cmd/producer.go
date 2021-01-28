package main

import (
	"config-manager/config"
	k "config-manager/infrastructure/kafka"
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	config := config.Get()
	ctx := context.Background()
	producer := k.StartProducer(config)

	fmt.Println("Producing message to ", producer.Topic)

	err := producer.WriteMessages(ctx, kafka.Message{
		Key:   []byte("new-connection"),
		Value: []byte(`{"hostname":"kafka.example","clientID":"9999"}`),
	})

	fmt.Println(err)
}
