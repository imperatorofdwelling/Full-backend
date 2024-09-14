package config

import (
	configDb "github.com/imperatorofdwelling/Website-backend/internal/config/db"
	configSrv "github.com/imperatorofdwelling/Website-backend/internal/config/server"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	DB     configDb.Database
	Server configSrv.Server
}

func LoadConfig() *Config {

	err := loadDotEnv("")
	if err != nil {
		log.Fatal(err)
	}

	srv := configSrv.InitServerConfig()

	db := configDb.InitDbConfig()

	return &Config{
		DB:     db,
		Server: srv,
	}
}

func loadDotEnv(filePath string) error {
	if filePath == "" {
		filePath = ".env"
	}
	err := godotenv.Load(filePath)
	return err
}
