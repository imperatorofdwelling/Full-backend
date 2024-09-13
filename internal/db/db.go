package db

import (
	"database/sql"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	config "github.com/imperatorofdwelling/Website-backend/internal/config/db"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

// ConnectToBD Подключение к PostgresSql по app.env
func ConnectToBD(cfg config.DataBase) (*sql.DB, error) {
	// Формирование строки подключения из конфига
	addr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.PsqlUser,
		cfg.PsqlPass,
		cfg.PsqlHost,
		cfg.PsqlPort,
		cfg.PsqlDBName,
	)
	psqlInfo := addr

	// Подключение к БД
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

//type Storage struct {
//	sync.Mutex
//	DB *sql.DB
//}
//
//var currStorage *Storage
//
//// InitDB initializes a new instance of the Database.
//func InitDB(cfg *Config) error {
//	// This if ensures that only 1 database instance is initialized.
//	if currStorage != nil {
//		return errors.New("the database is already initialized")
//	}
//	if cfg == nil {
//		return errors.New("config is empty")
//	}
//
//	connectionString := fmt.Sprintf(
//		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
//		cfg.DBUsername,
//		cfg.DBPassword,
//		cfg.DBHost,
//		cfg.DBPort,
//		cfg.DBName,
//		cfg.DBSSLMode,
//	)
//	db, err := sql.Open("postgres", connectionString)
//	if err != nil {
//		return err
//	}
//
//	if err = db.Ping(); err != nil {
//		return err
//	}
//
//	currStorage = &Storage{
//		DB: db,
//	}
//	return nil
//}
//
//func GetStorage() (*Storage, bool) {
//	if currStorage == nil {
//		return nil, false
//	}
//	return currStorage, true
//}
//
//func Disconnect() error {
//	storage, isContains := GetStorage()
//	if !isContains || storage.DB == nil {
//		return errors.New("the database is already initialized")
//	}
//	storage.Lock()
//	defer storage.Unlock()
//	err := storage.DB.Close()
//	return err
//}
