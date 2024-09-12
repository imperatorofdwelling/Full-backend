package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/imperatorofdwelling/Website-backend/internal/providers/user"
)

type ServerHTTP struct {
	Router *chi.Mux
}

func NewServerHTTP(userHandler *user.Handler) *ServerHTTP {
	r := chi.NewRouter()

	//r.Mount("", userHandler.NewHandler(r))

	return &ServerHTTP{Router: r}
}
