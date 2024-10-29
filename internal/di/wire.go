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
	ctrctProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/contracts"
	fvrtProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/favourite"
	flProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/file"
	locProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/location"
	resProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/reservation"
	srchProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/searchhistory"
	staysProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/stays"
	staysAdvProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/staysadvantage"
	staysReviewProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/staysreviews"
	usrProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/user"

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

		db.ConnectToBD,
		api.NewServerHTTP,
	))
}
