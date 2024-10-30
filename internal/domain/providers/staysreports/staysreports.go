package staysreports

import (
	"database/sql"
	"github.com/google/wire"
	staysReportHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/staysreports"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	staysReportRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/staysreports"
	staysReportSvc "github.com/imperatorofdwelling/Full-backend/internal/service/staysreports"
	"log/slog"
	"sync"
)

var (
	hdl     *staysReportHdl.Handler
	hdlOnce sync.Once

	svc     *staysReportSvc.Service
	svcOnce sync.Once

	repo     *staysReportRepo.Repo
	repoOnce sync.Once
)

var StaysReportsProvideSet wire.ProviderSet = wire.NewSet(
	ProvideStaysReportHandler,
	ProvideStaysReportService,
	ProvideStaysReportRepo,

	wire.Bind(new(interfaces.StaysReportsHandler), new(*staysReportHdl.Handler)),
	wire.Bind(new(interfaces.StaysReportsService), new(*staysReportSvc.Service)),
	wire.Bind(new(interfaces.StaysReportsRepo), new(*staysReportRepo.Repo)),
)

func ProvideStaysReportHandler(svc interfaces.StaysReportsService, log *slog.Logger) *staysReportHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &staysReportHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})
	return hdl
}

func ProvideStaysReportService(repo interfaces.StaysReportsRepo) *staysReportSvc.Service {
	svcOnce.Do(func() {
		svc = &staysReportSvc.Service{
			Repo: repo,
		}
	})
	return svc
}

func ProvideStaysReportRepo(db *sql.DB) *staysReportRepo.Repo {
	repoOnce.Do(func() {
		repo = &staysReportRepo.Repo{
			Db: db,
		}
	})
	return repo
}
