package database

import (
	"database/sql"

	"github.com/reynld/carbtographer/pkg/models"
)

// // GetItems returns all items per restuarant id
// func GetItems(id string, items *[]models.Items) {
// 	db.Where("r_id = ?", id).Find(&items)
// }

// // GetNames returns all restaurant names in database
// func GetNames(rest *[]models.Restaurants) {
// 	db.Find(&rest)
// }

// // GetIDByName returns the id of the restuarant
// func GetIDByName(name string, rest *[]models.Restaurants) {
// 	db.Table("restaurants").Where("name = ?", name).Find(&rest)
// }

// InsertRestaurant adds restaurant
func InsertRestaurant(db *sql.DB, name string, logo string, id *int) error {
	err := db.QueryRow(`INSERT INTO restaurant(name, logo)
		VALUES
		($1, $2)
		RETURNING id`, name, logo).Scan(id)
	if err != nil {
		return err
	}
	return nil
}

// InsertItem adds item with relation to restaurant
func InsertItem(db *sql.DB, item *models.JSONItem, restID int, id *int) error {
	err := db.QueryRow(`INSERT INTO 
		items(
			name, 
			type, 
			protein, 
			carbs, 
			fats, 
			calories, 
			alsperpro, 
			sodium, 
			rid
		)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`,
		item.Name,
		item.Type,
		item.Protein,
		item.Carbs,
		item.Fats,
		item.Calories,
		item.CalPerPro,
		item.Sodium,
		restID,
	).Scan(id)
	if err != nil {
		return err
	}
	return nil
}
