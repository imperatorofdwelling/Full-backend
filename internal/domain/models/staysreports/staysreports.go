package staysreports

import (
	"github.com/google/uuid"
	"time"
)

type StayReport struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	StayID       uuid.UUID `json:"stay_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	ReportAttach *string   `json:"report_attach,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type StaysReportEntity struct {
	ReportID    uuid.UUID
	UserName    string
	StayName    string
	Title       string
	Description string
}
