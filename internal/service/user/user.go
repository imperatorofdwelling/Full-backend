package user

import (
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
)

type Service struct {
	Repo interfaces.UserRepository
}
