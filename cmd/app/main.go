package main

import (
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/db"
	"github.com/imperatorofdwelling/Full-backend/internal/di"
	"log"
)

func main() {
	cfg := config.MustLoad()

	storage, err := db.NewStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	srv := di.Wire(storage.DB)

	srv.Start(cfg)
}
