package providers

import (
	"database/sql"
	"github.com/google/wire"
	staysHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/stays"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	staysRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/stays"
	staysSvc "github.com/imperatorofdwelling/Full-backend/internal/service/stays"
	"log/slog"
	"sync"
)

var (
	hdl     *staysHdl.Handler
	hdlOnce sync.Once

	svc     *staysSvc.Service
	svcOnce sync.Once

	repository     *staysRepo.Repo
	repositoryOnce sync.Once
)

var StaysProviderSet wire.ProviderSet = wire.NewSet(
	ProvideStaysHandler,
	ProvideStaysService,
	ProvideStaysRepo,

	wire.Bind(new(interfaces.StaysHandler), new(*staysHdl.Handler)),
	wire.Bind(new(interfaces.StaysService), new(*staysSvc.Service)),
	wire.Bind(new(interfaces.StaysRepo), new(*staysRepo.Repo)),
)

func ProvideStaysHandler(svc interfaces.StaysService, log *slog.Logger) *staysHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &staysHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideStaysService(repo interfaces.StaysRepo, locSvc interfaces.LocationService) *staysSvc.Service {
	svcOnce.Do(func() {
		svc = &staysSvc.Service{
			Repo:   repo,
			LocSvc: locSvc,
		}
	})

	return svc
}

func ProvideStaysRepo(db *sql.DB) *staysRepo.Repo {
	repositoryOnce.Do(func() {
		repository = &staysRepo.Repo{
			Db: db,
		}
	})

	return repository
}
