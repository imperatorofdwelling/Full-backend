//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"log/slog"
)

func InitializeAPI(cfg config.Config, log *slog.Logger) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectToBD, service.NewService, handler.NewHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
