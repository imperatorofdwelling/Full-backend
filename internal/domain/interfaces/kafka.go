package interfaces

import "github.com/imperatorofdwelling/Full-backend/internal/api/kafka"

type KafkaProducer interface {
	NewKafkaProducer() (*kafka.Producer, error)
	Close() error
	SendMessage(topic string, message []byte) error
}
