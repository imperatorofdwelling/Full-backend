package main

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/imperatorofdwelling/Website-backend/internal/config"
	db2 "github.com/imperatorofdwelling/Website-backend/internal/db"
	"log"
	"os"
)

func main() {
	cfg := config.LoadConfig()
	db, err := db2.ConnectToBD(cfg)
	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://cmd/migrator/migrations", "postgres", driver)
	if err != nil {
		panic(err)
	}

	cmd := os.Args[len(os.Args)-1]

	switch cmd {
	case "up":
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	default:
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}
}
