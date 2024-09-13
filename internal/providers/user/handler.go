package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/imperatorofdwelling/Full-backend/internal/domain"
	"net/http"
)

type Handler struct {
	svc domain.UserService
}

func (h *Handler) NewHandler(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/ping", h.FetchByUsername())
	})
}

func (h *Handler) FetchByUsername() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		render.JSON(w, r, map[string]string{
			"Hello": "World",
		})

		return
	}
}
