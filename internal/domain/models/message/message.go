package message

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID        uuid.UUID `json:"id"`
	ChatID    uuid.UUID `json:"chat_id"`
	UserID    uuid.UUID `json:"user_id"`
	Text      string    `json:"text"`
	Media     *string   `json:"media"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Entity struct {
	UserID    uuid.UUID `json:"user_id"`
	Text      string    `json:"text"`
	Media     *string   `json:"media"`
	UpdatedAt time.Time `json:"updated_at"`
}
