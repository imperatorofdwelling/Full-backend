package user

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Full-backend/internal/domain"
	"sync"
)

var (
	hdl     *Handler
	hdlOnce sync.Once

	svc     *service
	svcOnce sync.Once

	repo     *repository
	repoOnce sync.Once
)

var ProviderSet wire.ProviderSet = wire.NewSet(
	ProvideHandler,
	ProvideService,
	ProvideRepository,

	wire.Bind(new(domain.UserHandler), new(*Handler)),
	wire.Bind(new(domain.UserService), new(*service)),
	wire.Bind(new(domain.UserRepository), new(*repository)),
)

func ProvideHandler(svc domain.UserService) *Handler {
	hdlOnce.Do(func() {
		hdl = &Handler{
			svc: svc,
		}
	})

	return hdl
}

func ProvideService(repo domain.UserRepository) *service {
	svcOnce.Do(func() {
		svc = &service{
			repo: repo,
		}
	})

	return svc
}

func ProvideRepository(db *sql.DB) *repository {
	repoOnce.Do(func() {
		repo = &repository{
			db: db,
		}
	})

	return repo
}
