package newPassword

type NewPassword struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
} // @name NewPassword
