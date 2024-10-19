package user

import (
	"database/sql"
	"github.com/google/wire"
	fvrtHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/favourite"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	fvrtRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/favourite"
	fvrtSvc "github.com/imperatorofdwelling/Full-backend/internal/service/favourite"
	"log/slog"
	"sync"
)

var (
	hdl     *fvrtHdl.FavHandler
	hdlOnce sync.Once

	svc     *fvrtSvc.Service
	svcOnce sync.Once

	repo     *fvrtRepo.Repo
	repoOnce sync.Once
)

var FavouriteProviderSet = wire.NewSet(
	ProvideFavouriteHandler,
	ProvideFavouriteService,
	ProvideFavouriteRepository,

	wire.Bind(new(interfaces.FavouriteHandler), new(*fvrtHdl.FavHandler)),
	wire.Bind(new(interfaces.FavouriteService), new(*fvrtSvc.Service)),
	wire.Bind(new(interfaces.FavouriteRepo), new(*fvrtRepo.Repo)),
)

func ProvideFavouriteHandler(svc interfaces.FavouriteService, log *slog.Logger) *fvrtHdl.FavHandler {
	hdlOnce.Do(func() {
		hdl = &fvrtHdl.FavHandler{
			Svc: svc,
			Log: log,
		}
	})

	return hdl
}

func ProvideFavouriteService(repo interfaces.FavouriteRepo) *fvrtSvc.Service {
	svcOnce.Do(func() {
		svc = &fvrtSvc.Service{
			Repo: repo,
		}
	})

	return svc
}

func ProvideFavouriteRepository(db *sql.DB) *fvrtRepo.Repo {
	repoOnce.Do(func() {
		repo = &fvrtRepo.Repo{
			Db: db,
		}
	})

	return repo
}
