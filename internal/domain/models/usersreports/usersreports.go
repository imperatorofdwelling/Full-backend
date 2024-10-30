package usersreports

import (
	"github.com/google/uuid"
	"time"
)

type UsersReport struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	OwnerID      uuid.UUID `json:"owner_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	ReportAttach *string   `json:"report_attach,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UsersReportEntity struct {
	ReportID    uuid.UUID `json:"id"`
	UserName    string    `json:"user_name"`
	OwnerName   string    `json:"owner_name"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
