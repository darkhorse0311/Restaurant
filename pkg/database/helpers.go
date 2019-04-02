package database

import (
	"github.com/reynld/carbtographer/pkg/models"
)

// GetItems returns all items per restuarant id
func GetItems(id string, items *[]models.Items) {
	db.Where("r_id = ?", id).Find(&items)
}

// GetNames returns all restaurant names in database
func GetNames(rest *[]models.Restaurants) {
	db.Find(&rest)
}

// GetIDByName returns the id of the restuarant
func GetIDByName(name string, rest *[]models.Restaurants) {
	db.Table("restaurants").Where("name = ?", name).Find(&rest)
}
