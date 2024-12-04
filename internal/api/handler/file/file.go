package file

import (
	"github.com/go-chi/chi/v5"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	_ "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"log/slog"
	"net/http"
)

type Handler struct {
	Svc interfaces.FileService
	Log *slog.Logger
}

func (h *Handler) NewFileHandler(r chi.Router) {
	r.Route("/file", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Handle("/*", h.GetStaticImage())
		})
	})
}

func (h *Handler) GetStaticImage() http.Handler {
	fs := http.FileServer(http.Dir("static"))

	return http.Handler(http.StripPrefix("/api/v1/file/", fs))
}
