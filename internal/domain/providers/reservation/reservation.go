package reservation

import (
	"database/sql"
	"github.com/google/wire"
	resHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/reservation"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	resRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/reservation"
	resSvc "github.com/imperatorofdwelling/Full-backend/internal/service/reservation"
	"log/slog"
	"sync"
)

var (
	hdl     *resHdl.Handler
	hdlOnce sync.Once

	svc     *resSvc.Service
	svcOnce sync.Once

	repository     *resRepo.Repo
	repositoryOnce sync.Once
)

var ReservationProviderSet wire.ProviderSet = wire.NewSet(
	ProvideReservationHandler,
	ProvideReservationService,
	ProvideReservationRepository,

	wire.Bind(new(interfaces.ReservationHandler), new(*resHdl.Handler)),
	wire.Bind(new(interfaces.ReservationService), new(*resSvc.Service)),
	wire.Bind(new(interfaces.ReservationRepo), new(*resRepo.Repo)),
)

func ProvideReservationHandler(svc interfaces.ReservationService, log *slog.Logger) *resHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &resHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})
	return hdl
}

func ProvideReservationService(repo interfaces.ReservationRepo) *resSvc.Service {
	svcOnce.Do(func() {
		svc = &resSvc.Service{
			Repo: repo,
		}
	})

	return svc
}

func ProvideReservationRepository(db *sql.DB) *resRepo.Repo {
	repositoryOnce.Do(func() {
		repository = &resRepo.Repo{
			Db: db,
		}
	})

	return repository
}
