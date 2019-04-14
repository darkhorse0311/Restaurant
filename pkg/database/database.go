package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// getDBKeys makes sure all enviroment variables are set and return them
func getDBKeys() (map[string]string, error) {
	keys := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_NAME",
		"DB_PASSWORD",
		"ENVIROMENT",
	}

	values := map[string]string{}

	for _, key := range keys {
		v := os.Getenv(key)
		if v == "" {
			return nil, fmt.Errorf("eviroment variable %s is required", key)
		}
		values[key] = v
	}

	return values, nil
}

// getDBUri makes sure all enviroment variables are set and return them
func getDBUri() (string, error) {
	d, err := getDBKeys()
	if err != nil {
		return "", err
	}

	var dburi string
	if d["ENVIROMENT"] == "PRO" {
		// Production
		dburi = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s",
			d["DB_HOST"],
			d["DB_PORT"],
			d["DB_USER"],
			d["DB_NAME"],
			d["DB_PASSWORD"],
		)
	} else if d["ENVIROMENT"] == "DEV" {
		// Local
		dburi = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable",
			d["DB_HOST"],
			d["DB_PORT"],
			d["DB_USER"],
			d["DB_NAME"],
		)
	} else {
		return "", fmt.Errorf("ENVIROMENT variable must be PRO or DEV")
	}

	return dburi, nil
}

// InitializeDB connects to DB
func InitializeDB() *sql.DB {
	dburi, err := getDBUri()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", dburi)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
