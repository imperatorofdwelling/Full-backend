package config

import "os"

type DataBase struct {
	PsqlUser    string
	PsqlPass    string
	PsqlHost    string
	PsqlPort    string
	PsqlDBName  string
	PsqlSSLMode string
}

func (db *DataBase) Init() DataBase {
	db.PsqlUser = os.Getenv("POSTGRES_USER")
	db.PsqlPass = os.Getenv("POSTGRES_PASSWORD")
	db.PsqlHost = os.Getenv("POSTGRES_HOST")
	db.PsqlDBName = os.Getenv("POSTGRES_DB")
	db.PsqlSSLMode = os.Getenv("POSTGRES_DB_SSL")

	return *db
}
