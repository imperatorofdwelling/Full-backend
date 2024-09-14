package db

import (
	"database/sql"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/imperatorofdwelling/Website-backend/internal/config"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

func ConnectToBD(cfg *config.Config) (*sql.DB, error) {
	addr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.DbUser,
		cfg.DB.DbPass,
		cfg.DB.DbHost,
		cfg.DB.DbPort,
		cfg.DB.DbName,
	)
	psqlInfo := addr

	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
