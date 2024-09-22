package providers

import (
	"database/sql"
	"github.com/google/wire"
	locHdl "github.com/imperatorofdwelling/Website-backend/internal/api/handler/location"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/interfaces"
	locRepo "github.com/imperatorofdwelling/Website-backend/internal/repo/location"
	locSvc "github.com/imperatorofdwelling/Website-backend/internal/service/location"
	"log/slog"
	"sync"
)

var (
	hdl     *locHdl.Handler
	hdlOnce sync.Once

	svc     *locSvc.Service
	svcOnce sync.Once

	repository     *locRepo.Repo
	repositoryOnce sync.Once
)

var LocationProviderSet wire.ProviderSet = wire.NewSet(
	ProvideLocationHandler,
	ProvideLocationService,
	ProvideLocationRepository,

	wire.Bind(new(interfaces.LocationHandler), new(*locHdl.Handler)),
	wire.Bind(new(interfaces.LocationService), new(*locSvc.Service)),
	wire.Bind(new(interfaces.LocationRepo), new(*locRepo.Repo)),
)

func ProvideLocationHandler(svc interfaces.LocationService, log *slog.Logger) *locHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &locHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideLocationService(repo interfaces.LocationRepo) *locSvc.Service {
	svcOnce.Do(func() {
		svc = &locSvc.Service{
			Repo: repo,
		}
	})

	return svc
}

func ProvideLocationRepository(db *sql.DB) *locRepo.Repo {
	repositoryOnce.Do(func() {
		repository = &locRepo.Repo{
			Db: db,
		}
	})

	return repository
}
