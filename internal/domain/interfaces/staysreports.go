package interfaces

import (
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysreports"
	"golang.org/x/net/context"
	"net/http"
)

//go:generate mockery --name StaysReportsRepo
type StaysReportsRepo interface {
	CreateStaysReports(ctx context.Context, userId, stayId, title, description string) error
	GetAllStaysReports(ctx context.Context, userId string) ([]staysreports.StaysReportEntity, error)
	UpdateStaysReports(ctx context.Context, userId, stayId, title, description string) (*staysreports.StaysReportEntity, error)
	DeleteStaysReports(ctx context.Context, userId, reportId string) error
}

//go:generate mockery --name StaysReportsService
type StaysReportsService interface {
	CreateStaysReports(ctx context.Context, userId, stayId, title, description string) error
	GetAllStaysReports(ctx context.Context, userId string) ([]staysreports.StaysReportEntity, error)
	UpdateStaysReports(ctx context.Context, userId, stayId, title, description string) (*staysreports.StaysReportEntity, error)
	DeleteStaysReports(ctx context.Context, userId, reportId string) error
}

type StaysReportsHandler interface {
	CreateStaysReports(w http.ResponseWriter, r *http.Request)
	GetAllStaysReports(w http.ResponseWriter, r *http.Request)
	UpdateStaysReports(w http.ResponseWriter, r *http.Request)
	DeleteStaysReports(w http.ResponseWriter, r *http.Request)
}
