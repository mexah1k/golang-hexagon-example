package kafka

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/segmentio/kafka-go"
	"go.uber.org/fx"
	"golang-hexagon-example/internal/app/infrastructure/config"
)

func NewKafkaReader(config *config.AppConfig) *kafka.Reader {
	kafkaURL := config.Kafka.Broker
	if kafkaURL == "" && config.Environment == "Development" { // for local development
		kafkaURL = "localhost:9092"
	}

	topic := config.Kafka.Topic
	if topic == "" && config.Environment == "Development" { // for local development
		topic = "domain-demo"
	} else if topic == "" {
		log.Error("KAFKA_TOPIC is not set")
	}

	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
		MaxWait:  1e9,
		GroupID:  config.Kafka.GroupId,
	})
}

func registerKafkaConsumer(lifecycle fx.Lifecycle, kcs *ConsumerService) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return kcs.Start()
		},
		OnStop: func(ctx context.Context) error {
			return kcs.Stop()
		},
	})
}

var Module = fx.Options(
	fx.Provide(
		NewKafkaReader,
		NewKafkaConsumerService,
	),
	fx.Invoke(registerKafkaConsumer),
)
