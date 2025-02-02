package kafka

import (
	"context"
	"github.com/IBM/sarama"
	"log"
)

const (
	ConsumerGroup        = "payment-group"
	PaymentResponseTopic = "payment-response"
)

type Consumer struct {
	Group sarama.ConsumerGroup
}

func NewKafkaConsumer(groupHdl sarama.ConsumerGroupHandler) *Consumer {
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup(ServerAddr, ConsumerGroup, config)
	if err != nil {
		panic(err)
	}
	defer consumerGroup.Close()

	consumer := &Consumer{Group: consumerGroup}

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		consumer.Setup(ctx, consumerGroup, groupHdl)
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
