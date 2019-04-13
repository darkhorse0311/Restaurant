package models

//Items from databse
type Items struct {
	ID        int     `gorm:"primary_key" json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Protein   int     `json:"protein"`
	Carbs     int     `json:"carbs"`
	Fats      int     `json:"fats"`
	Calories  int     `json:"calories"`
	CalPerPro float32 `json:"calperpro"`
	Sodium    int     `json:"sodium"`
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
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	Photos   []string `json:"photos"`
	Distance float64  `json:"distance"`
	RID      int      `json:"r_id"`
}
