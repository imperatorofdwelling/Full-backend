//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Full-backend/internal/api"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/db"
	authProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/auth"
	locProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/location"
	usrProvider "github.com/imperatorofdwelling/Full-backend/internal/domain/providers/user"
	"log/slog"
)

func InitializeAPI(cfg *config.Config, log *slog.Logger) (*api.ServerHTTP, error) {
	panic(wire.Build(
		usrProvider.ProviderSet,
		locProvider.LocationProviderSet,
		authProvider.ProviderSet,

		db.ConnectToBD,
		api.NewServerHTTP,
	))
}
