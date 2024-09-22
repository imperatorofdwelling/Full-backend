package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	advHdl "github.com/imperatorofdwelling/Website-backend/internal/api/handler/advantage"
	locHdl "github.com/imperatorofdwelling/Website-backend/internal/api/handler/location"
	usrHdl "github.com/imperatorofdwelling/Website-backend/internal/api/handler/user"
	"github.com/imperatorofdwelling/Website-backend/internal/config"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
	"time"
)

type ServerHTTP struct {
	router *chi.Mux
}

func NewServerHTTP(
	cfg *config.Config,
	userHandler *usrHdl.UserHandler,
	locationHandler *locHdl.LocationHandler,
	advantageHandler *advHdl.Handler,
) *ServerHTTP {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(10 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
		userHandler.NewUserHandler(r)
		locationHandler.NewLocationHandler(r)
		advantageHandler.NewAdvantageHandler(r)
	})

	r.Get("/api/v1/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s:%s/api/v1/swagger/doc.json", cfg.Server.Addr, cfg.Server.Port)),
	))

	return &ServerHTTP{router: r}
}

func (sh *ServerHTTP) Start(cfg *config.Config, log *slog.Logger) {
	fmt.Print(fmt.Sprintf("Port is %s", cfg.Server.Port))
	log.Info(fmt.Sprintf("Starting server on port: %s", cfg.Server.Port))
	addr := cfg.Server.Addr + ":" + cfg.Server.Port
	err := http.ListenAndServe(addr, sh.router)
	if err != nil {
		log.Error(err.Error())
		return
	}
}
