package models

// JSONItem struct from json
type JSONItem struct {
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Protein   float32 `json:"protein"`
	Fats      float32 `json:"fats"`
	Carbs     float32 `json:"carbs"`
	Calories  float32 `json:"calories"`
	CalPerPro float32 `json:"calperpro"`
	Sodium    float32 `json:"Sodium"`
}

// JSONRestaurant struct from json
type JSONRestaurant struct {
	Name  string     `json:"name"`
	Logo  string     `json:"logo"`
	Items []JSONItem `json:"items"`
}

// Comu struct is the data passed between getInfo channels
type Comu struct {
	Items []JSONItem
	Index int
}
