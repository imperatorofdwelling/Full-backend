package providers

import (
	"database/sql"
	"github.com/google/wire"
	user3 "github.com/imperatorofdwelling/Website-backend/internal/api/handler/user"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/interfaces"
	user2 "github.com/imperatorofdwelling/Website-backend/internal/repo/user"
	"github.com/imperatorofdwelling/Website-backend/internal/service/user"
	"log/slog"
	"sync"
)

var (
	hdl     *user3.UserHandler
	hdlOnce sync.Once

	svc     *user.Service
	svcOnce sync.Once

	repo     *user2.Repository
	repoOnce sync.Once
)

var UserProviderSet wire.ProviderSet = wire.NewSet(
	ProvideUserHandler,
	ProvideUserService,
	ProvideUserRepository,

	wire.Bind(new(interfaces.UserHandler), new(*user3.UserHandler)),
	wire.Bind(new(interfaces.UserService), new(*user.Service)),
	wire.Bind(new(interfaces.UserRepository), new(*user2.Repository)),
)

func ProvideUserHandler(svc interfaces.UserService, log *slog.Logger) *user3.UserHandler {
	hdlOnce.Do(func() {
		hdl = &user3.UserHandler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideUserService(repo interfaces.UserRepository) *user.Service {
	svcOnce.Do(func() {
		svc = &user.Service{
			Repo: repo,
		}
	})

	return svc
}

func ProvideUserRepository(db *sql.DB) *user2.Repository {
	repoOnce.Do(func() {
		repo = &user2.Repository{
			Db: db,
		}
	})

	return repo
}
