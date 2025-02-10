package interfaces

import "net/http"

//go:generate mockery --name PaymentHandler
type PaymentHandler interface {
	MakePayment(w http.ResponseWriter, r *http.Request)
}
