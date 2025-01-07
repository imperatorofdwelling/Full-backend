package auth

import "github.com/imperatorofdwelling/Full-backend/pkg/validator"

type (
	Login struct {
		Email    string `json:"email" validate:"required" example:"user@example.com"`
		Password string `json:"password" validate:"required" example:"securepassword"`
		IsHashed bool   `json:"isHashed" validate:"required" example:"false"`
	} // @name Login

	Registration struct {
		Name     string `json:"name" validate:"required" example:"John Doe"`
		Email    string `json:"email" validate:"required" example:"user@example.com"`
		Password string `json:"password" validate:"required" example:"securepassword"`
		IsHashed bool   `json:"isHashed" validate:"required" example:"false"`
	} // @name Registration
)

func ValidateRegistration(v *validator.Validator, registration *Registration) {
	v.Check(registration.Name != "", "name", "Name field should not be empty")
	v.Check(len(registration.Name) > 2, "name", "length of the name must be greater than 2")

	v.Check(len(registration.Password) > 5, "password", "length of the password must be greater than 5")

	v.Check(validator.Matches(registration.Email, validator.EmailRX), "email", "must be in correct form")
}
