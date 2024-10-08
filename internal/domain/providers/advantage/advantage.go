package providers

import (
	"database/sql"
	"github.com/google/wire"
	advHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/advantage"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	advRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/advantage"
	advSvc "github.com/imperatorofdwelling/Full-backend/internal/service/advantage"
	"log/slog"
	"sync"
)

var (
	hdl     *advHdl.Handler
	hdlOnce sync.Once

	svc     *advSvc.Service
	svcOnce sync.Once

	repository     *advRepo.Repo
	repositoryOnce sync.Once
)

var AdvantageProviderSet wire.ProviderSet = wire.NewSet(
	ProvideAdvantageHandler,
	ProvideAdvantageService,
	ProvideAdvantageRepository,

	wire.Bind(new(interfaces.AdvantageHandler), new(*advHdl.Handler)),
	wire.Bind(new(interfaces.AdvantageService), new(*advSvc.Service)),
	wire.Bind(new(interfaces.AdvantageRepo), new(*advRepo.Repo)),
)

func ProvideAdvantageHandler(svc interfaces.AdvantageService, log *slog.Logger) *advHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &advHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})
	return hdl
}

func ProvideAdvantageService(repo interfaces.AdvantageRepo, fileSvc interfaces.FileService) *advSvc.Service {
	svcOnce.Do(func() {
		svc = &advSvc.Service{
			Repo:    repo,
			FileSvc: fileSvc,
		}
	})

	return svc
}

func ProvideAdvantageRepository(db *sql.DB) *advRepo.Repo {
	repositoryOnce.Do(func() {
		repository = &advRepo.Repo{
			Db: db,
		}
	})

	return repository
}
