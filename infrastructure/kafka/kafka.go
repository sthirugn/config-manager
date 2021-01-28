package kafka

import (
	"config-manager/application"
	"config-manager/domain"
	"context"
	"encoding/json"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

func NewResultsConsumer(cfg *viper.Viper) *kafka.Reader {
	consumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     cfg.GetStringSlice("KafkaBrokers"),
		Topic:       cfg.GetString("KafkaResultsTopic"),
		GroupID:     cfg.GetString("KafkaGroupID"),
		StartOffset: cfg.GetInt64("KafkaConsumerOffset"),
	})

	return consumer
}

func NewConnectionsConsumer(cfg *viper.Viper) *kafka.Reader {
	consumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     cfg.GetStringSlice("KafkaBrokers"),
		Topic:       cfg.GetString("KafkaConnectionsTopic"),
		GroupID:     cfg.GetString("KafkaGroupID"),
		StartOffset: cfg.GetInt64("KafkaConsumerOffset"),
	})

	return consumer
}

func StartReading(ctx context.Context, r *kafka.Reader, c *application.ConfigManagerService) {
	for {
		fmt.Println("Kafka reader - waiting on new connections")
		m, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("Reader - message from %s-%d [%d]: %s: %s\n",
			m.Topic,
			m.Partition,
			m.Offset,
			string(m.Key),
			string(m.Value))

		var newClients []domain.Client
		var client domain.Client
		if err := json.Unmarshal(m.Value, &client); err != nil {
			panic(err)
		}

		newClients = append(newClients, client)

		_, err = c.ApplyState("11111", "redhat", newClients)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func StartProducer(cfg *viper.Viper) *kafka.Writer {

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: cfg.GetStringSlice("KafkaBrokers"),
		Topic:   cfg.GetString("KafkaConnectionsTopic"),
	})

	return w
}
