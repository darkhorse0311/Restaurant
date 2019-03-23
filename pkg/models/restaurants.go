package models

// Restaurants {name: string}
type Restaurants struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
