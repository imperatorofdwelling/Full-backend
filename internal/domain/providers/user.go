package providers

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Website-backend/internal/api/handler"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/interfaces"
	repository "github.com/imperatorofdwelling/Website-backend/internal/repo"
	"github.com/imperatorofdwelling/Website-backend/internal/service"
	"log/slog"
	"sync"
)

var (
	hdl     *handler.UserHandler
	hdlOnce sync.Once

	svc     *service.UserService
	svcOnce sync.Once

	repo     *repository.UserRepository
	repoOnce sync.Once
)

var UserProviderSet wire.ProviderSet = wire.NewSet(
	ProvideUserHandler,
	ProvideUserService,
	ProvideUserRepository,

	wire.Bind(new(interfaces.UserHandler), new(*handler.UserHandler)),
	wire.Bind(new(interfaces.UserService), new(*service.UserService)),
	wire.Bind(new(interfaces.UserRepository), new(*repository.UserRepository)),
)

func ProvideUserHandler(svc interfaces.UserService, log *slog.Logger) *handler.UserHandler {
	hdlOnce.Do(func() {
		hdl = &handler.UserHandler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideUserService(repo interfaces.UserRepository) *service.UserService {
	svcOnce.Do(func() {
		svc = &service.UserService{
			Repo: repo,
		}
	})

	return svc
}

func ProvideUserRepository(db *sql.DB) *repository.UserRepository {
	repoOnce.Do(func() {
		repo = &repository.UserRepository{
			Db: db,
		}
	})

	return repo
}
