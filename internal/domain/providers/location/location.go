package providers

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Full-backend/internal/api/handler/location"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/repo/location"
	"github.com/imperatorofdwelling/Full-backend/internal/service/location"
	"log/slog"
	"sync"
)

var (
	hdl     *handler.LocationHandler
	hdlOnce sync.Once

	svc     *service.LocationService
	svcOnce sync.Once

	repository     *repo.LocationRepo
	repositoryOnce sync.Once
)

var LocationProviderSet wire.ProviderSet = wire.NewSet(
	ProvideLocationHandler,
	ProvideLocationService,
	ProvideLocationRepository,

	wire.Bind(new(interfaces.LocationHandler), new(*handler.LocationHandler)),
	wire.Bind(new(interfaces.LocationService), new(*service.LocationService)),
	wire.Bind(new(interfaces.LocationRepository), new(*repo.LocationRepo)),
)

func ProvideLocationHandler(svc interfaces.LocationService, log *slog.Logger) *handler.LocationHandler {
	hdlOnce.Do(func() {
		hdl = &handler.LocationHandler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideLocationService(repo interfaces.LocationRepository) *service.LocationService {
	svcOnce.Do(func() {
		svc = &service.LocationService{
			Repo: repo,
		}
	})

	return svc
}

func ProvideLocationRepository(db *sql.DB) *repo.LocationRepo {
	repositoryOnce.Do(func() {
		repository = &repo.LocationRepo{
			Db: db,
		}
	})

	return repository
}
