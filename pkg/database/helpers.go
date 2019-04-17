package database

import (
	"database/sql"

	"github.com/reynld/carbtographer/pkg/models"
)

// GetItems returns all items per restuarant id
func GetItems(db *sql.DB, id string) ([]models.Items, error) {
	rows, err := db.Query(`SELECT * FROM items WHERE r_id = $1`, id)
	if err != nil {
		return nil, err
	}

	items := []models.Items{}

	for rows.Next() {
		var item models.Items
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Type,
			&item.Protein,
			&item.Carbs,
			&item.Fats,
			&item.Calories,
			&item.CalPerPro,
			&item.Sodium,
			&item.RID,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// GetNames returns all restaurant names in database
func GetNames(db *sql.DB) ([]models.Restaurants, error) {
	rows, err := db.Query(`SELECT * FROM restaurants`)
	if err != nil {
		return nil, err
	}

	restuarants := []models.Restaurants{}
	for rows.Next() {
		var rest models.Restaurants
		err := rows.Scan(
			&rest.ID,
			&rest.Name,
			&rest.Logo,
		)
		if err != nil {
			return nil, err
		}

		restuarants = append(restuarants, rest)
	}

	return restuarants, nil
}

// GetRestaurantID returns the id of the restuarant
func GetRestaurantID(db *sql.DB, name string) (int, error) {
	var id int
	err := db.QueryRow(`SELECT id FROM restaurants WHERE name = $1`, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// InsertRestaurant adds restaurant
func InsertRestaurant(db *sql.DB, name string, logo string, id *int) error {
	err := db.QueryRow(`INSERT INTO restaurants(name, logo)
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
			calsperpro, 
			sodium, 
			r_id
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

// GetByUsername gets User by username
func GetByUsername(db *sql.DB, user *models.User, username string) error {
	err := db.QueryRow(
		`SELECT u.id, u.username, u.password FROM users u WHERE username = $1`,
		username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return err
	}
	return nil
}

// CreateUser returns User by username
func CreateUser(db *sql.DB, id *int, username string, password string) error {
	err := db.QueryRow(`INSERT INTO users(username, password)
		VALUES
		($1, $2)
		RETURNING id`, username, password).Scan(id)
	if err != nil {
		return err
	}
	return nil
}
