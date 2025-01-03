package chat

import (
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	ChatID      uuid.UUID  `json:"chat_id"`
	StayOwnerID uuid.UUID  `json:"stay_owner_id"`
	StayUserID  uuid.UUID  `json:"stay_user_id"`
	OperatorID  *uuid.UUID `json:"operator_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
} // @name Chat
