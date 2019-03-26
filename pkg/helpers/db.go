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
	host := os.Getenv("AWS_HOST")
	ap := os.Getenv("AWS_PORT")
	user := os.Getenv("AWS_USER")
	dbname := os.Getenv("AWS_DBNAME")
	pass := os.Getenv("AWS_PASSWORD")
	dburi := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, ap, user, dbname, pass)

	db, err := gorm.Open(
		"postgres",
		dburi)
	if err != nil {
		return nil, err
	}

	// Migrate and seed DB
	// migrateSeed(db)

	return db, nil
}

func migrateSeed(db *gorm.DB) {
	db.DropTableIfExists(&models.Restaurants{}, &models.Items{})
	db.AutoMigrate(&models.Restaurants{}, &models.Items{})

	// Path to json file to seed DB
	path := os.Getenv("SQLITE_PATH")
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
	jsonInfo := make([]Restuarant, 0)
	// Converst array of bytes to array of structs
	json.Unmarshal(byteValue, &jsonInfo)

	// Loops through array of resturatns to seed database
	for i, rest := range jsonInfo {
		db.Create(&models.Restaurants{Name: rest.Name})
		for _, item := range rest.Items {
			db.Create(&models.Items{
				Name:      item.Name,
				Type:      item.Type,
				Protein:   item.Protein,
				Carbs:     item.Carbs,
				Fats:      item.Fats,
				Calories:  item.Calories,
				CalPerPro: item.CalPerPro,
				Sodium:    item.Sodium,
				RID:       i + 1,
			})
		}
	}
}
