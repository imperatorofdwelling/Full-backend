package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
)

const (
	PaymentReqTopic = "payment"
	PayoutReqTopic  = "payout"
)

var ServerAddr = []string{"kafka-1:9090"}
//"kafka-2:9090",
//"kafka-3:9090",

type Producer struct {
	sync sarama.SyncProducer
}

func NewKafkaProducer() (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(ServerAddr, config)
	if err != nil {
		return nil, fmt.Errorf("error creating the Kafka producer: %s", err.Error())
	}

	return &Producer{
		sync: producer,
	}, nil
}

func (p *Producer) Close() error {
	return p.sync.Close()
}

func (p *Producer) SendMessage(topic string, key string, message interface{}) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshalling message: %s", err.Error())
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(messageBytes),
	}

	_, _, err = p.sync.SendMessage(msg)

	return err
}
