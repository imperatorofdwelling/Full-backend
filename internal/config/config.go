package config

import (
	"github.com/imperatorofdwelling/Website-backend/internal/db"
	"github.com/imperatorofdwelling/Website-backend/internal/di"
	"github.com/imperatorofdwelling/Website-backend/internal/server/http"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
)

type Config struct {
	SrvConfig *http.Config
	DBConfig  *db.Config
}

func LoadConfig(envFilePath string) *Config {
	err := loadDotEnv(envFilePath)
	if err != nil {
		log.Fatal(err)
	}
	serverConfig, err := http.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	dbConfig, err := db.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	cfg := &Config{
		SrvConfig: serverConfig,
		DBConfig:  dbConfig,
	}

	return cfg
}

func (c *Config) Run(sLog *slog.Logger) {
	err := db.InitDB(c.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	storage, isExist := db.GetStorage()
	if !isExist {
		log.Fatal("database not exist")
	}

	srvH := di.Wire(storage.DB)

	router := http.NewRouter(srvH.Router)

	srv := http.New(c.SrvConfig, router)

	defer c.Disconnect(srv)
	srv.Run()
}

func (c *Config) Disconnect(server *http.Server) {
	err := db.Disconnect()
	if err != nil {
		log.Println(err)
	}

	err = server.Disconnect()
	if err != nil {
		log.Println(err)
	}
}

func loadDotEnv(filePath string) error {
	if filePath == "" {
		filePath = ".env"
	}
	err := godotenv.Load(filePath)
	return err
}
