package auth

type (
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
