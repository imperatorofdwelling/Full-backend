package advantage

import (
	"github.com/gofrs/uuid"
	"time"
)

type (
	AdvantageEntity struct {
		Title string `json:"title,omitempty"`
		Image []byte `json:"image,omitempty"`
	} // @name AdvantageEntity

	Advantage struct {
		ID        uuid.UUID `json:"id"`
		Title     string    `json:"title"`
		Image     string    `json:"image,omitempty"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} // @name Advantage
)
