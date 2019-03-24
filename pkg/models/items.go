package models

//Items {}
type Items struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Protein  int    `json:"protein"`
	Carbs    int    `json:"carbs"`
	Fats     int    `json:"fats"`
	Calories int    `json:"calories"`
	Sodium   int    `json:"sodium"`
	RID      int    `sql:"type:integer REFERENCES restaurants(id) ON DELETE CASCADE ON UPDATE CASCADE" json:"r_id"`
}
