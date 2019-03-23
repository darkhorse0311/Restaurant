package models

//Items {}
type Items struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Protein  uint   `json:"protein"`
	Carbs    uint   `json:"carbs"`
	Fats     uint   `json:"fats"`
	Calories uint   `json:"calories"`
	Sodium   uint   `json:"sodium"`
	RID      uint   `sql:"type:integer REFERENCES restaurants(id) ON DELETE CASCADE ON UPDATE CASCADE" json:"r_id"`
}
