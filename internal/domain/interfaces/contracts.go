package interfaces

import (
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/contracts"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

//go:generate mockery --name ContractsRepo
type ContractsRepo interface {
	AddContract(ctx context.Context, userId, stayId string, dateStart, dateEnd time.Time) error
	UpdateContract(ctx context.Context, userId, stayId string, dateStart, dateEnd time.Time) (*contracts.ContractEntity, error)
	GetAllContracts(ctx context.Context, userId string) ([]contracts.ContractEntity, error)
}

//go:generate mockery --name ContractService
type ContractService interface {
	AddContract(ctx context.Context, userId, stayId string, dateStart, dateEnd time.Time) error
	UpdateContract(ctx context.Context, userId, stayId string, dateStart, dateEnd time.Time) (*contracts.ContractEntity, error)
	GetAllContracts(ctx context.Context, userId string) ([]contracts.ContractEntity, error)
}

type ContractHandler interface {
	AddContract(w http.ResponseWriter, r *http.Request)
	UpdateContract(w http.ResponseWriter, r *http.Request)
	GetAllContracts(w http.ResponseWriter, r *http.Request)
}
