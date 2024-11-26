package contracts

import (
	"database/sql"
	"github.com/google/wire"
	ctrctHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/contracts"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	ctrctRepo "github.com/imperatorofdwelling/Full-backend/internal/repo/contracts"
	ctrctSvc "github.com/imperatorofdwelling/Full-backend/internal/service/contracts"
	"log/slog"
	"sync"
)

var (
	contractHandler     *ctrctHdl.Handler
	contractHandlerOnce sync.Once

	contractRepo     *ctrctRepo.Repo
	contractRepoOnce sync.Once

	contractService     *ctrctSvc.Service
	contractServiceOnce sync.Once
)

var ContractProviderSet = wire.NewSet(
	ProvideContractHandler,
	ProvideContractService,
	ProvideContractRepository,

	wire.Bind(new(interfaces.ContractHandler), new(*ctrctHdl.Handler)),
	wire.Bind(new(interfaces.ContractService), new(*ctrctSvc.Service)),
	wire.Bind(new(interfaces.ContractsRepo), new(*ctrctRepo.Repo)),
)

func ProvideContractHandler(svc interfaces.ContractService, log *slog.Logger) *ctrctHdl.Handler {
	contractHandlerOnce.Do(func() {
		contractHandler = &ctrctHdl.Handler{
			Svc: svc,
			Log: log,
		}
	})
	return contractHandler
}

func ProvideContractService(repo interfaces.ContractsRepo) *ctrctSvc.Service {
	contractServiceOnce.Do(func() {
		contractService = &ctrctSvc.Service{
			Repo: repo,
		}
	})
	return contractService
}

func ProvideContractRepository(db *sql.DB) *ctrctRepo.Repo {
	contractRepoOnce.Do(func() {
		contractRepo = &ctrctRepo.Repo{
			Db: db,
		}
	})
	return contractRepo
}
