//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Website-backend/internal/api"
	"github.com/imperatorofdwelling/Website-backend/internal/config"
	"github.com/imperatorofdwelling/Website-backend/internal/db"
	locProvider "github.com/imperatorofdwelling/Website-backend/internal/domain/providers/location"
	usrProvider "github.com/imperatorofdwelling/Website-backend/internal/domain/providers/user"
	"log/slog"
)

func InitializeAPI(cfg *config.Config, log *slog.Logger) (*api.ServerHTTP, error) {
	panic(wire.Build(
		usrProvider.UserProviderSet,
		locProvider.LocationProviderSet,

		db.ConnectToBD,
		api.NewServerHTTP,
	))
}
