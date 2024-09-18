package user

import (
	"github.com/gofrs/uuid"
	"time"
)

type (
	Entity struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Phone     string    `json:"phone"`
		BirthDate time.Time `json:"birth_date,omitempty"`
		National  string    `json:"national,omitempty"`
		Gender    string    `json:"gender,omitempty"`
		Country   string    `json:"country,omitempty"`
		City      string    `json:"city,omitempty"`
		CreatedAt time.Time `json:"createdAt,omitempty"`
		UpdatedAt time.Time `json:"updatedAt,omitempty"`
	}

	User struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Phone     string    `json:"phone"`
		BirthDate string    `json:"birth_date"`
		National  string    `json:"national"`
		Gender    string    `json:"gender"`
		Country   string    `json:"country"`
		City      string    `json:"city"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}

	Info struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		BirthDate string `json:"birthDate"`
		National  string `json:"national"`
		Gender    string `json:"gender"`
	}

	Login struct {
		Email    string `json:"email" validate:"required, email"`
		Password string `json:"password" validate:"required" `
	}

	Registration struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required" `
	}
)
