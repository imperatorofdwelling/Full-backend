package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
)

const (
	PaymentTopic = "payment"
)

var ServerAddr = []string{"localhost:9094", "localhost:9095", "localhost:9096"}

type Producer struct {
	Sarama sarama.SyncProducer
}

func (p *Producer) NewKafkaProducer() (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(ServerAddr, config)
	if err != nil {
		return nil, fmt.Errorf("error creating the Kafka producer: %s", err.Error())
	}

	return &Producer{
		Sarama: producer,
	}, nil
}

func (p *Producer) Close() error {
	return p.Sarama.Close()
}

func (p *Producer) SendMessage(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	_, _, err := p.Sarama.SendMessage(msg)
	return err
}
