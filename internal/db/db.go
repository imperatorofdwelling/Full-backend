package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"sync"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func LoadConfig() (*Config, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	sslMode := os.Getenv("POSTGRES_DB_SSL")

	return &Config{
		DBHost:     host,
		DBPort:     port,
		DBUsername: user,
		DBPassword: pass,
		DBName:     dbName,
		DBSSLMode:  sslMode,
	}, nil
}

type Storage struct {
	sync.Mutex
	DB *sql.DB
}

var currStorage *Storage

// InitDB initializes a new instance of the Database.
func InitDB(cfg *Config) error {
	// This if ensures that only 1 database instance is initialized.
	if currStorage != nil {
		return errors.New("the database is already initialized")
	}
	if cfg == nil {
		return errors.New("config is empty")
	}

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMode,
	)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	currStorage = &Storage{
		DB: db,
	}
	return nil
}

func GetStorage() (*Storage, bool) {
	if currStorage == nil {
		return nil, false
	}
	return currStorage, true
}

func Disconnect() error {
	storage, isContains := GetStorage()
	if !isContains || storage.DB == nil {
		return errors.New("the database is already initialized")
	}
	storage.Lock()
	defer storage.Unlock()
	err := storage.DB.Close()
	return err
}
