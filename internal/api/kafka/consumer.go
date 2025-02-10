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

func NewKafkaConsumer(paymentConsumer *consumer.PaymentConsumerHdl) *Consumer {
	config := sarama.NewConfig()

	paymentConsumerHdl := paymentConsumer.NewPaymentConsumerHdl()

	consumerGroup, err := sarama.NewConsumerGroup(ServerAddr, ConsumerGroup, config)
	if err != nil {
		panic(err)
	}

	con := &Consumer{Group: consumerGroup}

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		con.Setup(ctx, consumerGroup, paymentConsumerHdl)
	}()

	return con
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
