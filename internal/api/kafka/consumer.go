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

	paymentConsumerHdl := paymentConsumer.NewPaymentConsumer()

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

//func (c *Consumer) SubscribeToResponse(key string, responseChan chan<- yoomodel.Payment) {
//	c.Group
//	partitionConsumer, err := p.Consumer.ConsumePartition(ResponseTopic, 0, sarama.OffsetNewest)
//	if err != nil {
//		p.Log.Error("failed to subscribe to response topic", "error", err.Error())
//		return
//	}
//	defer partitionConsumer.Close()
//
//	for {
//		select {
//		case msg := <-partitionConsumer.Messages():
//			var response Response
//			err := json.Unmarshal(msg.Value, &response)
//			if err != nil {
//				p.Log.Error("error unmarshalling response", "error", err.Error())
//				continue
//			}
//
//			// Проверяем, что ответ соответствует нашему запросу
//			if string(msg.Key) == key {
//				responseChan <- response
//				return // Завершаем подписку после получения ответа
//			}
//		case err := <-partitionConsumer.Errors():
//			p.Log.Error("error consuming response", "error", err.Error())
//			return
//		}
//	}
//}
