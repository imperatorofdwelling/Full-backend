package main

import (
	_ "github.com/imperatorofdwelling/Full-backend/docs"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/di"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
)

// @title IOD App API
// @version 1.0
// description API Server for IOD application

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @externalDocs.description OpenAPI
// @externalDocs.url https://swagger.io/resources/open-api/
func main() {
	cfg := config.LoadConfig()
	log := logger.New(logger.EnvLocal)

	if server, err := di.InitializeAPI(cfg, log); err == nil {
		server.Start(cfg, log)
	}
}
