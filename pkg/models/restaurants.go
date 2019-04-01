package models

// Restaurants resposne from databse
type Restaurants struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}
