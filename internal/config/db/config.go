package config

import "os"

type Database struct {
	DbUser string `env-required:"true"`
	DbPass string `env-required:"true"`
	DbHost string `env-required:"true"`
	DbPort string `env-required:"true"`
	DbName string `env-required:"true"`
}

func InitDbConfig() Database {
	return Database{
		DbUser: os.Getenv("POSTGRES_USER"),
		DbPass: os.Getenv("POSTGRES_PASSWORD"),
		DbHost: os.Getenv("POSTGRES_HOST"),
		DbPort: os.Getenv("POSTGRES_PORT"),
		DbName: os.Getenv("POSTGRES_DB"),
	}
}

//import "os"
//
//type DataBase struct {
//	PsqlUser    string
//	PsqlPass    string
//	PsqlHost    string
//	PsqlPort    string
//	PsqlDBName  string
//	PsqlSSLMode string
//}
//
//func InitDbConfig() *DataBase {
//	return &DataBase{
//		PsqlUser:    os.Getenv("POSTGRES_USER"),
//		PsqlPass:    os.Getenv("POSTGRES_PASSWORD"),
//		PsqlHost:    os.Getenv("POSTGRES_HOST"),
//		PsqlPort:    os.Getenv("POSTGRES_PORT"),
//		PsqlDBName:  os.Getenv("POSTGRES_DB"),
//		PsqlSSLMode: "disable",
//	}
//}
