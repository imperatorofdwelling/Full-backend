package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"log/slog"
)

type UserHandler struct {
	Svc interfaces.UserService
	Log *slog.Logger
}

func (h *UserHandler) NewUserHandler(r chi.Router) {
	r.Route("/user", func(r chi.Router) {

	})
}
