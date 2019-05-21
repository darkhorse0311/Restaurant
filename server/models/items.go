package models

import (
	"database/sql"
)

//Items from databse
type Items struct {
	ID        int     `gorm:"primary_key" json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Protein   float32 `json:"protein"`
	Carbs     float32 `json:"carbs"`
	Fats      float32 `json:"fats"`
	Calories  float32 `json:"calories"`
	CalPerPro float32 `json:"calperpro"`
	Sodium    float32 `json:"sodium"`
	RID       int     `sql:"type:integer REFERENCES restaurants(id) ON DELETE CASCADE ON UPDATE CASCADE" json:"r_id"`
}

// Restaurants resposne from databse
type Restaurants struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

// Business db struct
type Business struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Coordinates struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
	} `json:"coordinates"`
	Photos   []string `json:"photos"`
	Distance float32  `json:"distance"`
	RID      int      `json:"r_id"`
}

// GetItems returns all items per restuarant id
func GetItems(db *sql.DB, id string) ([]Items, error) {
	rows, err := db.Query(`SELECT * FROM items WHERE r_id = $1`, id)
	if err != nil {
		return nil, err
	}

	items := []Items{}

	for rows.Next() {
		var item Items
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
func GetNames(db *sql.DB) ([]Restaurants, error) {
	rows, err := db.Query(`SELECT * FROM restaurants`)
	if err != nil {
		return nil, err
	}

	restuarants := []Restaurants{}
	for rows.Next() {
		var rest Restaurants
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
func InsertItem(db *sql.DB, item *JSONItem, restID int, id *int) error {
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
