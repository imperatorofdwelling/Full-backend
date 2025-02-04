package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka/consumer"
	"log"
)

const (
	ConsumerGroup        = "payment-group"
	PaymentResponseTopic = "payment-response"
)

type Consumer struct {
	Group sarama.ConsumerGroup
}

func NewKafkaConsumer(paymentConsumer *consumer.PaymentConsumer) *Consumer {
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup(ServerAddr, ConsumerGroup, config)
	if err != nil {
		panic(err)
	}
	defer consumerGroup.Close()

	consumer := &Consumer{Group: consumerGroup}

	paymentConsumerHdl := paymentConsumer.NewPaymentConsumer()

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		consumer.Setup(ctx, consumerGroup, paymentConsumerHdl)
	}()

	return consumer
}

func (c *Consumer) Setup(ctx context.Context, group sarama.ConsumerGroup, hdl sarama.ConsumerGroupHandler) error {
	for {
		err := group.Consume(ctx, []string{PaymentReqTopic, PayoutReqTopic, PaymentResponseTopic}, hdl)
		if err != nil {
			log.Printf("Error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return nil
		}
	}
}
