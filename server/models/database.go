package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"golang.org/x/crypto/bcrypt"
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

// RunSeeds migrates and seeds databse with JSON file from scraper
func RunSeeds(db *sql.DB) {
	// Path to json file to seed DB
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "restaurantData.json")
	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// Convert file to array of bytes
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// Makes a array of Restuarant strucs
	jsonInfo := make([]JSONRestaurant, 0)
	// Converst array of bytes to array of structs
	json.Unmarshal(byteValue, &jsonInfo)

	// Loops through array of resturatns to seed database
	for i, rest := range jsonInfo {
		fmt.Printf("INSERTING: %d:%d \n", i+1, len(jsonInfo))

		var id int
		err := InsertRestaurant(db, rest.Name, rest.Logo, &id)
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range rest.Items {
			var iID int

			err := InsertItem(db, &item, id, &iID)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte("pass"), 10)
	if err != nil {
		log.Fatal("error hasing seed password")
	}

	_, err = CreateUser(db, "rey", string(hash))
	if err != nil {
		log.Fatal("error seeding user")
	}
}

// RunMigrations runs migrations on database
func RunMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
