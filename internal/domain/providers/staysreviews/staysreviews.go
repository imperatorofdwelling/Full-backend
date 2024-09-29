package providers

import (
	"database/sql"
	"github.com/google/wire"
	stRevHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/staysreviews"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	stRevRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/staysreviews"
	stRevSvc "github.com/imperatorofdwelling/Full-backend/internal/service/staysreviews"
	"log/slog"
	"sync"
)

var (
	hdl     *stRevHdl.Handler
	hdlOnce sync.Once

	svc     *stRevSvc.Service
	svcOnce sync.Once

	repository     *stRevRepo.Repo
	repositoryOnce sync.Once
)

var StaysReviewsProviderSet wire.ProviderSet = wire.NewSet(
	ProvideStaysReviewsHandler,
	ProvideStaysReviewsService,
	ProvideStaysReviewsRepository,

	wire.Bind(new(interfaces.StaysReviewsHandler), new(*stRevHdl.Handler)),
	wire.Bind(new(interfaces.StaysReviewsService), new(*stRevSvc.Service)),
	wire.Bind(new(interfaces.StaysReviewsRepo), new(*stRevRepo.Repo)),
)

func ProvideStaysReviewsHandler(svc interfaces.StaysReviewsService, log *slog.Logger) *stRevHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &stRevHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})
	return hdl
}

func ProvideStaysReviewsService(repo interfaces.StaysReviewsRepo) *stRevSvc.Service {
	svcOnce.Do(func() {
		svc = &stRevSvc.Service{
			Repo: repo,
		}
	})

	return svc
}

func ProvideStaysReviewsRepository(db *sql.DB) *stRevRepo.Repo {
	repositoryOnce.Do(func() {
		repository = &stRevRepo.Repo{
			Db: db,
		}
	})

	return repository
}
