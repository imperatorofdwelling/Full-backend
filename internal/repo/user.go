package repo

import (
	"context"
	"database/sql"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/models"
)

type UserRepository struct {
	Db *sql.DB
}

func (r *UserRepository) FetchByUsername(ctx context.Context, username string) (*models.UserEntity, error) {
	panic("implement me")
}
