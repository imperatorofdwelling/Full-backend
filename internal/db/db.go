package db

import (
	"database/sql"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"sync"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Storage struct {
	sync.Mutex
	DB *sql.DB
}

func NewStorage(cfg *config.Config) (*Storage, error) {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DB.DbUser,
		cfg.DB.DbPass,
		cfg.DB.DbHost,
		cfg.DB.DbPort,
		cfg.DB.DbName,
		"disable",
	)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{
		DB: db,
	}, nil
}
