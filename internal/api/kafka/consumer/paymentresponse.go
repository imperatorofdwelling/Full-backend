package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	yoomodel "github.com/eclipsemode/go-yookassa-sdk/yookassa/model"
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
	const op = "kafka.consumer.paymentresponse.ConsumeClaim"
	for msg := range claim.Messages() {

		requestID := string(msg.Key)

		var payment yoomodel.Payment

		if err := json.Unmarshal(msg.Value, &payment); err != nil {
			return fmt.Errorf("%s: %s", op, err.Error())
		}

		if payment.Status == "" {
			return fmt.Errorf("%s: %s", op, "internal status error")
		}

		if responseChan, ok := c.WaitForResponse[requestID]; ok {
			responseChan <- PaymentResponse{
				RequestID: requestID,
				Result:    payment,
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
