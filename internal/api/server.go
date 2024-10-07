package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	advHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/advantage"
	authHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/auth"
	locHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/location"
	reservationHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/reservation"
	staysHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/stays"
	staysAdvHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/staysadvantage"
	usrHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/user"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
	"time"
)

type ServerHTTP struct {
	router http.Handler
}

func NewServerHTTP(
	cfg *config.Config,
	authHandler *authHdl.AuthHandler,
	userHandler *usrHdl.UserHandler,
	locationHandler *locHdl.Handler,
	advantageHandler *advHdl.Handler,
	staysHandler *staysHdl.Handler,
	staysAdvHandler *staysAdvHdl.Handler,
	reservationHandler *reservationHdl.Handler,
) *ServerHTTP {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Route("/api/v1/", func(r chi.Router) {
		authHandler.NewAuthHandler(r)
		staysAdvHandler.NewStaysAdvantageHandler(r)
		reservationHandler.NewReservationHandler(r)
		staysHandler.NewStaysHandler(r)
	})
	// Маршруты защищенные JWTMiddleware
	r.Group(func(r chi.Router) {
		r.Use(authHandler.JWTMiddleware)
		userHandler.NewUserHandler(r)
		locationHandler.NewLocationHandler(r)
		advantageHandler.NewAdvantageHandler(r)

	})

	r.Get("/api/v1/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s:%s/api/v1/swagger/doc.json", "localhost", cfg.Server.Port)),
	))

	handler := cors.Default().Handler(r)

	return &ServerHTTP{router: handler}
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
