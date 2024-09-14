//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Website-backend/internal/api"
	"github.com/imperatorofdwelling/Website-backend/internal/config"
	"github.com/imperatorofdwelling/Website-backend/internal/db"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/providers"
	"log/slog"
)

func InitializeAPI(cfg *config.Config, log *slog.Logger) (*api.ServerHTTP, error) {
	panic(wire.Build(
		providers.UserProviderSet,

		db.ConnectToBD,
		api.NewServerHTTP,
	))
}
