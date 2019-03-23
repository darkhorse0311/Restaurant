package models

import "github.com/jinzhu/gorm"

// Restaurants {name: string}
type Restaurants struct {
	gorm.Model
	Name string
}
