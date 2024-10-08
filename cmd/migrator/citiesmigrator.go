package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Location struct {
	City            string  `json:"city"`
	FederalDistrict string  `json:"federal_district"`
	FiasId          string  `json:"fias_id"`
	KladrId         string  `json:"kladr_id"`
	Lat             string  `json:"lat"`
	Lon             string  `json:"lon"`
	Okato           int     `json:"okato"`
	Oktmo           int     `json:"oktmo"`
	Population      float64 `json:"population"`
	RegionIsoCode   string  `json:"region_iso_code"`
	RegionName      string  `json:"region_name"`
}

func main() {

	dbPath, err := getDbPath()
	if err != nil {
		log.Fatal(fmt.Errorf("error getting database path: %v", err))
	}

	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		log.Fatal(fmt.Errorf("error opening database connection: %v", err))
	}

	isExists, err := checkMigrationDone(db)
	if err != nil {
		panic(err)
	}
	if isExists {
		fmt.Println("migration already done")
		return
	}

	content, err := os.ReadFile("cities.json")
	if err != nil {
		panic(err)
	}

	var locations []Location

	err = json.Unmarshal(content, &locations)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO locations (city, federal_district, fias_id, kladr_id, lat, lon, okato, oktmo, population, region_iso_code, region_name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	if err != nil {
		log.Fatal(fmt.Errorf("error preparing statement: %v", err))
	}

	for _, location := range locations {
		_, err := stmt.Exec(location.City, location.FederalDistrict, location.FiasId, location.KladrId, location.Lat, location.Lon, strconv.Itoa(location.Okato), strconv.Itoa(location.Oktmo), location.Population, location.RegionIsoCode, location.RegionName)
		if err != nil {
			log.Fatal(fmt.Errorf("error inserting row: %v", err))
		}
	}

	fmt.Println("Data has been successfully inserted into the table!")
}

func checkMigrationDone(db *sql.DB) (bool, error) {
	stmt, err := db.Prepare("SELECT EXISTS(SELECT 1 FROM locations)")
	if err != nil {
		return false, fmt.Errorf("error preparing statement: %v", err)
	}

	var exists bool

	err = stmt.QueryRow().Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking migration existence: %v", err)
	}

	return exists, nil
}

func getDbPath() (string, error) {
	var dbUserName, dbPass, dbHost, dbName string

	flag.StringVar(&dbUserName, "db-user-name", "postgres", "Name of the database user")
	flag.StringVar(&dbPass, "db-pass", "", "Password of the database user")
	flag.StringVar(&dbHost, "db-host", "localhost:5432", "Port of the database")
	flag.StringVar(&dbName, "db-name", "", "Name of the database to migrate")
	flag.Parse()

	if dbName == "" {
		return "", fmt.Errorf("db-name is required")
	}

	dbPath := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUserName, dbPass, dbHost, dbName)

	return dbPath, nil
}
