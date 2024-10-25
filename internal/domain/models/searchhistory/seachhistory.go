package searchhistory

import (
	"github.com/google/uuid"
	"time"
)

type SearchHistoryEntity struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type SearchHistory struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
