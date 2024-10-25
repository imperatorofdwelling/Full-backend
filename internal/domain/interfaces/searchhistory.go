package interfaces

import (
	"context"
	history "github.com/imperatorofdwelling/Full-backend/internal/domain/models/searchhistory"
	"net/http"
)

type SearchHistoryRepo interface {
	AddHistory(ctx context.Context, userId, name string) error
	GetAllHistoryByUserId(ctx context.Context, userId string) ([]history.SearchHistory, error)
}

type SearchHistoryService interface {
	AddHistory(ctx context.Context, userId, name string) error
	GetAllHistoryByUserId(ctx context.Context, userId string) ([]history.SearchHistory, error)
}

type SearchHistoryHandler interface {
	AddHistory(w http.ResponseWriter, r *http.Request)
	GetAllHistoryByUserId(w http.ResponseWriter, r *http.Request)
}
