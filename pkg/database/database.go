package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/reynld/carbtographer/pkg/models"
)

var db *gorm.DB
var err error

// migrateSeed migrates and seeds databse with JSON file from scraper
func migrateSeed() {
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
	jsonInfo := make([]restuarant, 0)
	// Converst array of bytes to array of structs
	json.Unmarshal(byteValue, &jsonInfo)

	// Loops through array of resturatns to seed database
	for i, rest := range jsonInfo {
		fmt.Printf("inserting: %d:%d \n", i, len(jsonInfo))
		db.Create(&models.Restaurants{Name: rest.Name, Logo: rest.Logo})
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

// getDBKeys makes sure all enviroment variables are set and return them
func getDBKeys() (map[string]string, error) {
	keys := []string{
		"AWS_HOST",
		"AWS_PORT",
		"AWS_USER",
		"AWS_DBNAME",
		"AWS_PASSWORD",
		"YELP_API_KEY",
	}

	d := map[string]string{}

	for _, key := range keys {
		v := os.Getenv(key)
		if v == "" {
			return nil, fmt.Errorf("eviroment variable %s is required", key)
		}
		d[key] = v
	}

	return d, nil
}

// InitDB creates, migrates and seeds database
func InitDB() (*gorm.DB, error) {

	k, err := getDBKeys()
	if err != nil {
		panic(err)
	}

	// Production
	dburi := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", k["AWS_HOST"], k["AWS_PORT"], k["AWS_USER"], k["AWS_DBNAME"], k["AWS_PASSWORD"])

	// Local
	// dburi := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", k["AWS_HOST"], k["AWS_PORT"], k["AWS_USER"], k["AWS_DBNAME"])

	db, err = gorm.Open("postgres", dburi)
	if err != nil {
		return nil, err
	}
	// Migrate and seed DB
	// migrateSeed()
	return db, nil
}
