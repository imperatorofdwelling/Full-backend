//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Website-backend/internal/api"
	"github.com/imperatorofdwelling/Website-backend/internal/config"
	"github.com/imperatorofdwelling/Website-backend/internal/db"
	advProvider "github.com/imperatorofdwelling/Website-backend/internal/domain/providers/advantage"
	flProvider "github.com/imperatorofdwelling/Website-backend/internal/domain/providers/file"
	locProvider "github.com/imperatorofdwelling/Website-backend/internal/domain/providers/location"
	staysProvider "github.com/imperatorofdwelling/Website-backend/internal/domain/providers/stays"
	usrProvider "github.com/imperatorofdwelling/Website-backend/internal/domain/providers/user"
	"log/slog"
)

func InitializeAPI(cfg *config.Config, log *slog.Logger) (*api.ServerHTTP, error) {
	panic(wire.Build(
		usrProvider.UserProviderSet,
		locProvider.LocationProviderSet,
		advProvider.AdvantageProviderSet,
		flProvider.FileProviderSet,
		staysProvider.StaysProviderSet,

		db.ConnectToBD,
		api.NewServerHTTP,
	))
}
