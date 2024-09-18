package service

import (
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
)

type UserService struct {
	Repo interfaces.UserRepository
}
