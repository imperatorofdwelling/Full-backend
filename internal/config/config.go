package config

import (
	"errors"
	"flag"
	"fmt"
	configDb "github.com/imperatorofdwelling/Full-backend/internal/config/db"
	configSrv "github.com/imperatorofdwelling/Full-backend/internal/config/server"
	"github.com/joho/godotenv"
	"log"
)

type Environment string

const (
	LocalEnv Environment = "local"
	DevEnv   Environment = "dev"
	ProdEnv  Environment = "prod"
)

var GlobalEnv Environment

type Config struct {
	DB     configDb.Database
	Server configSrv.Server
}

func LoadConfig() *Config {
	err := loadFlags()
	if err != nil {
		log.Fatal(err)
	}

	err = loadDotEnv()
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

func loadDotEnv() error {
	filePath := fmt.Sprintf(".env.%s", GlobalEnv)

	err := godotenv.Load(filePath)
	return err
}

func loadFlags() error {
	envFlag := flag.String("env", string(LocalEnv), "Environment type")
	flag.Parse()

	switch Environment(*envFlag) {
	case LocalEnv, DevEnv, ProdEnv:
		GlobalEnv = Environment(*envFlag)
	default:
		return errors.New("invalid environment type")
	}

	return nil
}
