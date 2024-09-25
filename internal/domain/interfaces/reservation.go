package interfaces

//go:generate mockery --name ReservationRepo
type ReservationRepo interface {
}

//go:generate mockery --name ReservationService
type ReservationService interface {
}

type ReservationHandler interface {
}
