package staysadvantage

import (
	"database/sql"
	"github.com/google/wire"
	advHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/staysadvantage"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	advRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/staysadvantage"
	advSvc "github.com/imperatorofdwelling/Full-backend/internal/service/staysadvantage"
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

var StaysAdvantageProviderSet wire.ProviderSet = wire.NewSet(
	ProvideStaysAdvantageHandler,
	ProvideStaysAdvantageService,
	ProvideStaysAdvantageRepo,

	wire.Bind(new(interfaces.StaysAdvantageHandler), new(*advHdl.Handler)),
	wire.Bind(new(interfaces.StaysAdvantageService), new(*advSvc.Service)),
	wire.Bind(new(interfaces.StaysAdvantageRepo), new(*advRepo.Repo)),
)

func ProvideStaysAdvantageHandler(svc interfaces.StaysAdvantageService, log *slog.Logger) *advHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &advHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideStaysAdvantageService(repo interfaces.StaysAdvantageRepo, staysSvc interfaces.StaysService, advantageSvc interfaces.AdvantageService) *advSvc.Service {
	svcOnce.Do(func() {
		svc = &advSvc.Service{
			Repo:    repo,
			StaySvc: staysSvc,
			AdvSvc:  advantageSvc,
		}
	})

	return svc
}

func ProvideStaysAdvantageRepo(db *sql.DB) *advRepo.Repo {
	repositoryOnce.Do(func() {
		repository = &advRepo.Repo{
			Db: db,
		}
	})

	return repository
}
