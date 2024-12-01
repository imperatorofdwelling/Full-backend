package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	advHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/advantage"
	authHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/auth"
	chatHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/chat"
	ctrctHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/contracts"
	fvrtHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/favourite"
	imgHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/image"
	locHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/location"
	msgHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/message"
	reservationHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/reservation"
	srchRevHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/searchhistory"
	staysHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/stays"
	staysAdvHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/staysadvantage"
	staysReportRevHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/staysreports"
	staysRevHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/staysreviews"
	usrHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/user"
	usersReportHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/usersreports"
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
	staysReviewsHandler *staysRevHdl.Handler,
	favouriteHandler *fvrtHdl.FavHandler,
	searchHandler *srchRevHdl.Handler,
	contractHandler *ctrctHdl.Handler,
	staysReportHandler *staysReportRevHdl.Handler,
	usersReportHandler *usersReportHdl.Handler,
	messageHandler *msgHdl.Handler,
	chatHandler *chatHdl.Handler,
	imageHandler *imgHdl.Handler,
) *ServerHTTP {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Route("/api/v1/", func(r chi.Router) {
		authHandler.NewAuthHandler(r)
		advantageHandler.NewAdvantageHandler(r)
		staysAdvHandler.NewStaysAdvantageHandler(r)
		userHandler.NewUserHandler(r)
		locationHandler.NewLocationHandler(r)
		reservationHandler.NewReservationHandler(r)
		staysReviewsHandler.NewStaysReviewsHandler(r)
		staysHandler.NewStaysHandler(r)
		favouriteHandler.NewFavouriteHandler(r)
		searchHandler.NewHistorySearchHandler(r)
		contractHandler.NewContractHandler(r)
		staysReportHandler.NewStaysReportsHandler(r)
		usersReportHandler.NewUsersReportsHandler(r)
		messageHandler.NewMessageHandler(r)
		chatHandler.NewChatHandler(r)
		imageHandler.NewImageHandler(r)

		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL("http://81.200.153.83/api/v1/swagger/doc.json"),
		))
	})

	// TODO Change CORS in production
	handler := cors.AllowAll().Handler(r)

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
