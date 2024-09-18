package auth

import (
	"database/sql"
	"github.com/google/wire"
	authHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/auth"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	authRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/auth"
	authSvc "github.com/imperatorofdwelling/Full-backend/internal/service/auth"
	"log/slog"
	"sync"
)

var (
	hdl     *authHdl.AuthHandler
	hdlOnce sync.Once

	svc     *authSvc.Service
	svcOnce sync.Once

	repo     *authRepo.Repository
	repoOnce sync.Once
)

var ProviderSet wire.ProviderSet = wire.NewSet(
	ProvideAuthHandler,
	ProvideAuthService,
	ProvideAuthRepository,

	wire.Bind(new(interfaces.AuthHandler), new(*authHdl.AuthHandler)),
	wire.Bind(new(interfaces.AuthService), new(*authSvc.Service)),
	wire.Bind(new(interfaces.AuthRepository), new(*authRepo.Repository)),
)

func ProvideAuthHandler(svc interfaces.AuthService, log *slog.Logger) *authHdl.AuthHandler {
	hdlOnce.Do(func() {
		hdl = &authHdl.AuthHandler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideAuthService(authRepo interfaces.AuthRepository, userRepo interfaces.UserRepository) *authSvc.Service {
	svcOnce.Do(func() {
		svc = &authSvc.Service{
			AuthRepo: authRepo,
			UserRepo: userRepo,
		}
	})

	return svc
}

func ProvideAuthRepository(db *sql.DB) *authRepo.Repository {
	repoOnce.Do(func() {
		repo = &authRepo.Repository{
			Db: db,
		}
	})

	return repo
}
