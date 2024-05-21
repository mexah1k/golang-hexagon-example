package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"golang-hexagon-example/internal/app/core/ports"
	"log"
	"strings"
)

// ConsumerService manages Kafka message consumption.
type ConsumerService struct {
	reader *kafka.Reader
	dcs    ports.UrlService
}

// NewKafkaConsumerService creates a new instance of KafkaConsumerService.
func NewKafkaConsumerService(reader *kafka.Reader, dcs ports.UrlService) *ConsumerService {
	return &ConsumerService{reader: reader, dcs: dcs}
}

// Start begins the message consumption from Kafka.
func (k *ConsumerService) Start() error {
	log.Println("Starting Kafka consumer...")
	go func() {
		for {
			m, err := k.reader.ReadMessage(context.Background())
			if err != nil {
				log.Printf("Failed to read messages: %v", err)
				continue
			}
			log.Println("Received message: ", string(m.Key))
			urls := strings.Split(string(m.Value), ",") // Assuming CSV format
			if err := k.dcs.Analyze(urls); err != nil {
				log.Printf("Failed to process domains: %v", err)
			}
		}
	}()
	return nil
}

func (k *ConsumerService) Stop() error {
	log.Println("Stopping Kafka consumer...")
	return k.reader.Close()
}
