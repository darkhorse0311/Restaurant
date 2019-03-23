package helpers

import (
	"github.com/jinzhu/gorm"
	"github.com/reynld/carbtographer/pkg/models"
)

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
	// Create
	db.Create(&models.Restaurants{Name: "Res 1"})
	db.Create(&models.Restaurants{Name: "Res 2"})
	db.Create(&models.Items{Name: "Item 1", Type: "", Protein: 1, Carbs: 2, Fats: 3, Calories: 4, Sodium: 5, RID: 1})
}
