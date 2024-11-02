package interfaces

import (
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/usersreports"
	"golang.org/x/net/context"
	"net/http"
)

//go:generate mockery --name UsersReportsRepo
type UsersReportsRepo interface {
	CreateUsersReports(ctx context.Context, userId, toBlameId, title, description string) error
	GetAllUsersReports(ctx context.Context, userId string) ([]usersreports.UsersReportEntity, error)
	UpdateUsersReports(ctx context.Context, userId, toBlameId, title, description string) (*usersreports.UsersReportEntity, error)
	DeleteUsersReports(ctx context.Context, userId, reportId string) error
}

//go:generate mockery --name UsersReportsService
type UsersReportsService interface {
	CreateUsersReports(ctx context.Context, userId, toBlameId, title, description string) error
	GetAllUsersReports(ctx context.Context, userId string) ([]usersreports.UsersReportEntity, error)
	UpdateUsersReports(ctx context.Context, userId, toBlameId, title, description string) (*usersreports.UsersReportEntity, error)
	DeleteUsersReports(ctx context.Context, userId, reportId string) error
}

type UsersReportsHandler interface {
	CreateUsersReports(w http.ResponseWriter, r *http.Request)
	GetAllUsersReports(w http.ResponseWriter, r *http.Request)
	UpdateUsersReports(w http.ResponseWriter, r *http.Request)
	DeleteUsersReports(w http.ResponseWriter, r *http.Request)
}
