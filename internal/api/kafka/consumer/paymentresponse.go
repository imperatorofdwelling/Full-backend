package consumer

import (
	"github.com/IBM/sarama"
	"log/slog"
)

type PaymentConsumerHdl struct {
	Log             *slog.Logger
	WaitForResponse map[string]chan PaymentResponse
}

type PaymentResponse struct {
	RequestID string
	Result    interface{}
}

func (*PaymentConsumerHdl) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*PaymentConsumerHdl) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (c *PaymentConsumerHdl) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {

		requestID := string(msg.Key)
		value := string(msg.Value)

		if responseChan, ok := c.WaitForResponse[requestID]; ok {
			responseChan <- PaymentResponse{
				RequestID: requestID,
				Result:    value,
			}
		}

		sess.MarkMessage(msg, "")
	}

	return nil
}

func (c *PaymentConsumerHdl) NewPaymentConsumerHdl() *PaymentConsumerHdl {

	return &PaymentConsumerHdl{
		Log:             c.Log,
		WaitForResponse: c.WaitForResponse,
	}
}
