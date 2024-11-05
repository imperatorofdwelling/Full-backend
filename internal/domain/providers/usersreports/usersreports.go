package usersreports

import (
	"database/sql"
	"github.com/google/wire"
	usersReportHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/usersreports"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	usersReportRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/usersreports"
	usersReportSvc "github.com/imperatorofdwelling/Full-backend/internal/service/usersreports"
	"log/slog"
	"sync"
)

var (
	hdl     *usersReportHdl.Handler
	hdlOnce sync.Once

	svc     *usersReportSvc.Service
	svcOnce sync.Once

	repo     *usersReportRepo.Repo
	repoOnce sync.Once
)

var UsersReportsProvideSet wire.ProviderSet = wire.NewSet(
	ProvideUsersReportHandler,
	ProvideUsersReportService,
	ProvideUsersReportRepo,

	wire.Bind(new(interfaces.UsersReportsHandler), new(*usersReportHdl.Handler)),
	wire.Bind(new(interfaces.UsersReportsService), new(*usersReportSvc.Service)),
	wire.Bind(new(interfaces.UsersReportsRepo), new(*usersReportRepo.Repo)),
)

func ProvideUsersReportHandler(svc interfaces.UsersReportsService, log *slog.Logger) *usersReportHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &usersReportHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})
	return hdl
}

func ProvideUsersReportService(repo interfaces.UsersReportsRepo) *usersReportSvc.Service {
	svcOnce.Do(func() {
		svc = &usersReportSvc.Service{
			Repo: repo,
		}
	})
	return svc
}

func ProvideUsersReportRepo(db *sql.DB) *usersReportRepo.Repo {
	repoOnce.Do(func() {
		repo = &usersReportRepo.Repo{
			Db: db,
		}
	})
	return repo
}
