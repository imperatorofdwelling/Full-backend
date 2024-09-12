//go:build wireinject

package di

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Website-backend/internal/api"
	"github.com/imperatorofdwelling/Website-backend/internal/providers/user"
)

func Wire(db *sql.DB) *api.ServerHTTP {
	panic(wire.Build(user.ProviderSet, api.NewServerHTTP))
}
