package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/reynld/carbtographer/pkg/models"
)

// Items struct from json
type Items struct {
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Protein   int     `json:"protein"`
	Fats      int     `json:"fats"`
	Carbs     int     `json:"carbs"`
	Calories  int     `json:"calories"`
	CalPerPro float32 `json:"calperpro"`
	Sodium    int     `json:"Sodium"`
}

// Restuarant struct from json
type Restuarant struct {
	Name  string  `json:"name"`
	Items []Items `json:"items"`
}

// InitDB creates, migrates and seeds database
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./carbtographer.sqlite3")
	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=reynaldo dbname=reynaldo")
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	db.AutoMigrate(&models.Restaurants{}, &models.Items{})
	// Create foreign key relation // Does not work on SQLite3
	// db.Model(&Items{}).AddForeignKey("r_id", "restaurants(id)", "RESTRICT", "RESTRICT")

	SeedDB(db)

	return db, nil
}

// SeedDB seeds database Restuarant and Items Tables
func SeedDB(db *gorm.DB) {

	// Open our jsonFile
	jsonFile, err := os.Open("/Users/reynaldo/go/src/github.com/reynld/carbtographer/restuarantData.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// Convert file to array of bytes
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Makes a array of Restuarant strucs
	jsonInfo := make([]Restuarant, 0)

	// Converst array of bytes to array of structs
	json.Unmarshal(byteValue, &jsonInfo)

	// Loops through array of resturatns to seed database
	for i, rest := range jsonInfo {
		db.Create(&models.Restaurants{Name: rest.Name})
		for _, item := range rest.Items {
			db.Create(&models.Items{
				Name:     item.Name,
				Type:     item.Type,
				Protein:  item.Protein,
				Carbs:    item.Carbs,
				Fats:     item.Fats,
				Calories: item.Calories,
				Sodium:   item.Sodium,
				RID:      i + 1,
			})
		}
	}
}
