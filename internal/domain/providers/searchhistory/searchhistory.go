package searchhistory

import (
	"database/sql"
	"github.com/google/wire"
	srchHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/searchhistory"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	srchRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/searchhistory"
	srchSvc "github.com/imperatorofdwelling/Full-backend/internal/service/searchhistory"
	"log/slog"
	"sync"
)

var (
	searchHistoryHandler     *srchHdl.Handler
	searchHistoryHandlerOnce sync.Once

	searchHistoryService     *srchSvc.Service
	searchHistoryServiceOnce sync.Once

	searchHistoryRepo     *srchRepo.Repo
	searchHistoryRepoOnce sync.Once
)

var SearchHistoryProviderSet = wire.NewSet(
	ProvideSearchHistoryHandler,
	ProvideSearchHistoryService,
	ProvideSearchHistoryRepository,

	wire.Bind(new(interfaces.SearchHistoryHandler), new(*srchHdl.Handler)),
	wire.Bind(new(interfaces.SearchHistoryService), new(*srchSvc.Service)),
	wire.Bind(new(interfaces.SearchHistoryRepo), new(*srchRepo.Repo)),
)

func ProvideSearchHistoryHandler(svc interfaces.SearchHistoryService, log *slog.Logger) *srchHdl.Handler {
	searchHistoryHandlerOnce.Do(func() {
		searchHistoryHandler = &srchHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})
	return searchHistoryHandler
}

func ProvideSearchHistoryService(repo interfaces.SearchHistoryRepo) *srchSvc.Service {
	searchHistoryServiceOnce.Do(func() {
		searchHistoryService = &srchSvc.Service{
			Repo: repo,
		}
	})
	return searchHistoryService
}

func ProvideSearchHistoryRepository(db *sql.DB) *srchRepo.Repo {
	searchHistoryRepoOnce.Do(func() {
		searchHistoryRepo = &srchRepo.Repo{
			Db: db,
		}
	})
	return searchHistoryRepo
}
