package models

type (
	UserEntity struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Phone     string `json:"phone"`
		BirthDate string `json:"birthDate"`
		National  string `json:"national,omitempty"`
		Gender    string `json:"gender,omitempty"`
	}

	User struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		BirthDate string `json:"birthDate"`
		National  string `json:"national"`
		Gender    string `json:"gender"`
	}
)
