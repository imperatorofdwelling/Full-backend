package providers

import (
	"database/sql"
	"github.com/google/wire"
	usrHdl "github.com/imperatorofdwelling/Website-backend/internal/api/handler/user"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/interfaces"
	usrRepo "github.com/imperatorofdwelling/Website-backend/internal/repo/user"
	usrSvc "github.com/imperatorofdwelling/Website-backend/internal/service/user"
	"log/slog"
	"sync"
)

var (
	hdl     *usrHdl.UserHandler
	hdlOnce sync.Once

	svc     *usrSvc.UserService
	svcOnce sync.Once

	repo     *usrRepo.UserRepository
	repoOnce sync.Once
)

var UserProviderSet wire.ProviderSet = wire.NewSet(
	ProvideUserHandler,
	ProvideUserService,
	ProvideUserRepository,

	wire.Bind(new(interfaces.UserHandler), new(*usrHdl.UserHandler)),
	wire.Bind(new(interfaces.UserService), new(*usrSvc.UserService)),
	wire.Bind(new(interfaces.UserRepository), new(*usrRepo.UserRepository)),
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

func ProvideUserService(repo interfaces.UserRepository) *usrSvc.UserService {
	svcOnce.Do(func() {
		svc = &usrSvc.UserService{
			Repo: repo,
		}
	})

	return svc
}

func ProvideUserRepository(db *sql.DB) *usrRepo.UserRepository {
	repoOnce.Do(func() {
		repo = &usrRepo.UserRepository{
			Db: db,
		}
	})

	return repo
}
