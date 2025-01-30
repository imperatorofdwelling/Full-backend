package user

import (
	"database/sql"
	"github.com/google/wire"
	usrHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/user"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	usrRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/user"
	usrSvc "github.com/imperatorofdwelling/Full-backend/internal/service/user"
	"log/slog"
	"sync"
)

var (
	hdl     *usrHdl.UserHandler
	hdlOnce sync.Once

	svc     *usrSvc.Service
	svcOnce sync.Once

	repo     *usrRepo.Repository
	repoOnce sync.Once
)

var ProviderSet wire.ProviderSet = wire.NewSet(
	ProvideUserHandler,
	ProvideUserService,
	ProvideUserRepository,

	wire.Bind(new(interfaces.UserHandler), new(*usrHdl.UserHandler)),
	wire.Bind(new(interfaces.UserService), new(*usrSvc.Service)),
	wire.Bind(new(interfaces.UserRepository), new(*usrRepo.Repository)),
)

func ProvideUserHandler(svc interfaces.UserService, log *slog.Logger) *usrHdl.UserHandler {
	hdlOnce.Do(func() {
		hdl = &usrHdl.UserHandler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideUserService(userRepo interfaces.UserRepository, fileSvc interfaces.FileService, confirmRepo interfaces.ConfirmEmailRepository) *usrSvc.Service {
	svcOnce.Do(func() {
		svc = &usrSvc.Service{
			UserRepo:         userRepo,
			ConfirmEmailRepo: confirmRepo,
			FileSvc:          fileSvc,
		}
	})

	return svc
}

func ProvideUserRepository(db *sql.DB) *usrRepo.Repository {
	repoOnce.Do(func() {
		repo = &usrRepo.Repository{
			Db: db,
		}
	})

	return repo
}
