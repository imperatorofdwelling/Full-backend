package config

import "os"

type Database struct {
	User         string `env-required:"true"`
	Password     string `env-required:"true"`
	Host         string `env-required:"true"`
	Port         string `env-required:"true"`
	DriverName   string `env-required:"true"`
	DatabaseName string `env-required:"true"`
	SSLMode      string `env-required:"true"`
}

func InitDbConfig() Database {
	return Database{
		User:         os.Getenv("POSTGRES_USER"),
		Password:     os.Getenv("POSTGRES_PASSWORD"),
		Host:         os.Getenv("POSTGRES_HOST"),
		Port:         os.Getenv("POSTGRES_PORT"),
		DatabaseName: os.Getenv("POSTGRES_DB"),
		DriverName:   os.Getenv("POSTGRES_DRIVER"),
		SSLMode:      os.Getenv("POSTGRES_SSL_MODE"),
	}
}
