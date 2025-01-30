package passwordOTP

type PasswordOTP struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp" validate:"required"`
} // @name PasswordOTP
