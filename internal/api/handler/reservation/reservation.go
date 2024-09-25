package reservation

import (
	"github.com/go-chi/chi/v5"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"log/slog"
)

type Handler struct {
	Svc interfaces.ReservationService
	Log *slog.Logger
}

func (h *Handler) NewReservationHandler(r chi.Router) {
	r.Route("/reservation", func(r chi.Router) {

	})
}
