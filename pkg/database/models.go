package database

// items struct from json
type items struct {
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Protein   int     `json:"protein"`
	Fats      int     `json:"fats"`
	Carbs     int     `json:"carbs"`
	Calories  int     `json:"calories"`
	CalPerPro float32 `json:"calperpro"`
	Sodium    int     `json:"Sodium"`
}

// restuarant struct from json
type restuarant struct {
	Name  string  `json:"name"`
	Logo  string  `json:"logo"`
	Items []items `json:"items"`
}
