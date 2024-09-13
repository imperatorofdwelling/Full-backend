package account

type Account struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	BirthDate string `json:"birthDate"`
	National  string `json:"national"`
	Gender    string `json:"gender"`
}

type Info struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	BirthDate string `json:"birthDate"`
	National  string `json:"national"`
	Gender    string `json:"gender"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required" `
}

type Registration struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required" `
}

//type (
//	User struct {
//		ID       string `json:"id"`
//		Username string `json:"username"`
//	}
//
//	UserEntity struct {
//		ID       string
//		Username string
//		Password string
//	}
//
//	UserRepository interface {
//		FetchByUsername(ctx context.Context, username string) (*UserEntity, error)
//	}
//
//	UserService interface {
//		FetchByUsername(ctx context.Context, username string) (*User, error)
//	}
//
//	UserHandler interface {
//		FetchByUsername() http.HandlerFunc
//	}
//)
