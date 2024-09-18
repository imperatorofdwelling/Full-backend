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

	svc     *authSvc.AuthService
	svcOnce sync.Once

	repo     *authRepo.AuthRepository
	repoOnce sync.Once
)

var AuthProviderSet wire.ProviderSet = wire.NewSet(
	ProvideAuthHandler,
	ProvideAuthService,
	ProvideAuthRepository,

	wire.Bind(new(interfaces.AuthHandler), new(*authHdl.AuthHandler)),
	wire.Bind(new(interfaces.AuthService), new(*authSvc.AuthService)),
	wire.Bind(new(interfaces.AuthRepository), new(*authRepo.AuthRepository)),
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

func ProvideAuthService(authRepo interfaces.AuthRepository, userRepo interfaces.UserRepository) *authSvc.AuthService {
	svcOnce.Do(func() {
		svc = &authSvc.AuthService{
			AuthRepo: authRepo,
			UserRepo: userRepo,
		}
	})

	return svc
}

func ProvideAuthRepository(db *sql.DB) *authRepo.AuthRepository {
	repoOnce.Do(func() {
		repo = &authRepo.AuthRepository{
			Db: db,
		}
	})

	return repo
}
