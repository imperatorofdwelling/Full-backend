package consumer

import (
	"fmt"
	"github.com/IBM/sarama"
	"log/slog"
)

type PaymentConsumer struct {
	Log *slog.Logger
}

func (*PaymentConsumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*PaymentConsumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (consumer *PaymentConsumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Println("Idempotency Key", string(msg.Key))
		fmt.Println("Hello WOOOOORLD: ", string(msg.Value))

		sess.MarkMessage(msg, "")
	}
	return nil
}

func (*PaymentConsumer) NewPaymentConsumer() *PaymentConsumer {
	return &PaymentConsumer{}
}
