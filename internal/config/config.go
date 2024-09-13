package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Server struct {
	Addr         string        `env-default:"localhost"`
	Port         string        `env-default:"8080"`
	ReadTimeout  time.Duration `env-default:"5s"`
	WriteTimeout time.Duration `env-default:"10s"`
	IdleTimeout  time.Duration `env-default:"60s"`
}

type Database struct {
	DbUser string `env-required:"true"`
	DbPass string `env-required:"true"`
	DbHost string `env-required:"true"`
	DbPort string `env-required:"true"`
	DbName string `env-required:"true"`
}

type Config struct {
	DB     Database
	Server Server
}

func MustLoad() *Config {
	err := loadDotEnv()
	if err != nil {
		log.Fatal(err)
	}

	srv := Server{
		Addr: os.Getenv("SERVER_ADDR"),
		Port: os.Getenv("SERVER_PORT"),
	}

	db := Database{
		DbUser: os.Getenv("POSTGRES_USER"),
		DbPass: os.Getenv("POSTGRES_PASSWORD"),
		DbHost: os.Getenv("POSTGRES_HOST"),
		DbPort: os.Getenv("POSTGRES_PORT"),
		DbName: os.Getenv("POSTGRES_DB"),
	}

	return &Config{
		DB:     db,
		Server: srv,
	}
}

func loadDotEnv() error {
	err := godotenv.Load()
	return err
}
