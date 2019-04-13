package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/reynld/carbtographer/pkg/models"
)

// RunSeeds migrates and seeds databse with JSON file from scraper
func RunSeeds(db *sql.DB) {
	// Path to json file to seed DB
	path := os.Getenv("JSON_PATH")
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
	jsonInfo := make([]models.JSONRestaurant, 0)
	// Converst array of bytes to array of structs
	json.Unmarshal(byteValue, &jsonInfo)

	// Loops through array of resturatns to seed database
	for i, rest := range jsonInfo {
		fmt.Printf("INSERTING: %d:%d \n", i, len(jsonInfo))

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
}
