//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Full-backend/internal/api"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/db"
	advProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/advantage"
	authProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/auth"
	chatProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/chat"
	confirmEmailProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/confirmEmail"
	ctrctProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/contracts"
	fvrtProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/favourite"
	flProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/file"
	kafkaProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/kafka"
	locProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/location"
	msgProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/message"
	paymentProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/payment"
	resProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/reservation"
	srchProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/searchhistory"
	staysProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/stays"
	staysAdvProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/staysadvantage"
	staysReportsProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/staysreports"
	staysReviewProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/staysreviews"
	usrProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/user"
	usrsReportsProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/usersreports"

	"log/slog"
)

func InitializeAPI(cfg *config.Config, log *slog.Logger) (*api.ServerHTTP, error) {
	panic(wire.Build(
		usrProvider.ProviderSet,
		locProvider.LocationProviderSet,
		authProvider.ProviderSet,
		advProvider.AdvantageProviderSet,
		flProvider.FileProviderSet,
		staysProvider.StaysProviderSet,
		staysAdvProvider.StaysAdvantageProviderSet,
		resProvider.ReservationProviderSet,
		staysReviewProvider.StaysReviewsProviderSet,
		fvrtProvider.FavouriteProviderSet,
		srchProvider.SearchHistoryProviderSet,
		ctrctProvider.ContractProviderSet,
		staysReportsProvider.StaysReportsProvideSet,
		usrsReportsProvider.UsersReportsProvideSet,
		msgProvider.MessageProviderSet,
		chatProvider.ChatProviderSet,
		confirmEmailProvider.ProvideSet,
		kafkaProvider.KafkaProviderSet,
		paymentProvider.PaymentProviderSet,

		db.ConnectToBD,
		api.NewServerHTTP,
	))
}
