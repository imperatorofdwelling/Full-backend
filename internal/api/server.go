package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	advHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/advantage"
	authHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/auth"
	chatHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/chat"
	confirmEmailHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/confirmEmail"
	ctrctHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/contracts"
	fvrtHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/favourite"
	fileHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/file"
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
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka"
	paymentHdl "github.com/imperatorofdwelling/Full-backend/internal/api/kafka/payment"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
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
	fileHandler *fileHdl.Handler,
	confirmEmailHandler *confirmEmailHdl.Handler,
	kafkaProducer *kafka.Producer,
	paymentHandler *paymentHdl.Handler,
) *ServerHTTP {
	r := chi.NewRouter()

	kafkaProducer, err := kafkaProducer.NewKafkaProducer()
	if err != nil {
		log.Fatal("error creating kafka producer")
	}
	defer kafkaProducer.Close()

	r.Use(middleware.RequestID)
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
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
		fileHandler.NewFileHandler(r)
		confirmEmailHandler.NewConfirmEmailHandler(r)
		paymentHandler.NewPaymentHandler(r)

		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("http://%s:%s/api/v1/swagger/doc.json", cfg.Server.Host, cfg.Server.Port)),
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
