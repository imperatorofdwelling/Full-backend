package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/interfaces"
	"net/http"
)

type UserHandler struct {
	Svc interfaces.UserService
}

func (h *UserHandler) NewUserHandler(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/ping", h.FetchByUsername())
	})
}

func (h *UserHandler) FetchByUsername() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		render.JSON(w, r, map[string]string{
			"Hello": "World",
		})

		return
	}
}
