package interfaces

import (
	"context"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/account"
	"log/slog"
)

type ServiceUseCase interface {
	Migrate(ctx context.Context, log *slog.Logger) error

	//Account
	Registration(ctx context.Context, newAccount account.Registration) (*account.Info, error)
	Login(ctx context.Context, acc account.Login) (int64, error)
	PutAccount(ctx context.Context, id string, updateAcc account.Info) (*account.Info, error)
}
