package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/providers/user"
	"log"
	"net/http"
)

type ServerHTTP struct {
	router *chi.Mux
}

func NewServerHTTP(userHandler *user.Handler) *ServerHTTP {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		userHandler.NewHandler(r)
	})

	return &ServerHTTP{router: r}
}

func (sh *ServerHTTP) Start(cfg *config.Config) {
	log.Print("Server starting...")
	addr := fmt.Sprintf("%s:%s", cfg.Server.Addr, cfg.Server.Port)
	err := http.ListenAndServe(addr, sh.router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
