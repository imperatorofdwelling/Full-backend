package main

import (
	"github.com/imperatorofdwelling/Website-backend/internal/config"
	"github.com/imperatorofdwelling/Website-backend/internal/di"
	"github.com/imperatorofdwelling/Website-backend/pkg/logger"
)

func main() {
	cfg := config.LoadConfig()
	log := logger.New(logger.EnvLocal)

	if server, err := di.InitializeAPI(cfg, log); err == nil {
		server.Start(cfg, log)
	}
}
