package config

import (
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

// import "os"
//
//	type Host struct {
//		IP   string
//		Port string
//	}
func InitServerConfig() Server {
	return Server{
		Addr: os.Getenv("SERVER_ADDR"),
		Port: os.Getenv("SERVER_PORT"),
	}
}
