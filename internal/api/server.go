package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Website-backend/internal/api/handler"
	config "github.com/imperatorofdwelling/Website-backend/internal/config/server"
	"log/slog"
	"net/http"
	"time"
)

type ServerHTTP struct {
	router *chi.Mux
}

func NewServerHTTP(userHandler *handler.Handler, log *slog.Logger) *ServerHTTP {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Get("/login", userHandler.Login(log))
	r.Post("/registration", userHandler.Registration(log))
	r.Use(handler.JWTMiddleware(log))
	r.Route("/accounts", userHandler.Account(log))
	//r.Get("/protected", userHandler.Protected)
	//r.Route("/accounts", userHandler.Account)

	return &ServerHTTP{router: r}
}

func (sh *ServerHTTP) Start(cfg config.Host, log *slog.Logger) {
	log.Info("Server starting")
	addr := cfg.IP + ":" + cfg.Port
	err := http.ListenAndServe(addr, sh.router)
	if err != nil {
		log.Error(err.Error())
		return
	}
}
