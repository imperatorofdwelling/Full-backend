package main

import (
	"github.com/imperatorofdwelling/Website-backend/internal/config"
	"github.com/imperatorofdwelling/Website-backend/internal/di"
	"github.com/imperatorofdwelling/Website-backend/pkg/logger"
)

func main() {
	cfg := config.LoadConfig()
	log := logger.New(logger.EnvLocal)

	server, diErr := di.InitializeAPI(cfg.Db, log)
	if diErr != nil {
		log.Error("cannot start server: ", diErr)
	} else {
		server.Start(cfg.Host, log)
	}
}
