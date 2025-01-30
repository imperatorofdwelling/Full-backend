package searchhistory

import (
	"github.com/gofrs/uuid"
	"time"
)

type (
	SearchHistoryEntity struct {
		ID        uuid.UUID `json:"id"`
		UserID    uuid.UUID `json:"user_id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} // @name SearchHistoryEntity

	SearchHistory struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} // @name SearchHistory
)
