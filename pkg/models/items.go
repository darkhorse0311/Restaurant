package models

import "github.com/jinzhu/gorm"

//Items {}
type Items struct {
	gorm.Model
	Name     string
	Type     string
	Protein  uint
	Carbs    uint
	Fats     uint
	Calories uint
	Sodium   uint
	RID      uint `sql:"type:integer REFERENCES restaurants(id) ON DELETE CASCADE ON UPDATE CASCADE"`
}
