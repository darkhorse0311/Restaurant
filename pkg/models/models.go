package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

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

// JSONItem struct from json
type JSONItem struct {
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Protein   int     `json:"protein"`
	Fats      int     `json:"fats"`
	Carbs     int     `json:"carbs"`
	Calories  int     `json:"calories"`
	CalPerPro float32 `json:"calperpro"`
	Sodium    int     `json:"Sodium"`
}

// JSONRestaurant struct from json
type JSONRestaurant struct {
	Name  string     `json:"name"`
	Logo  string     `json:"logo"`
	Items []JSONItem `json:"items"`
}

// Claims a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
	jwt.StandardClaims
}

// JWTResponse struct returned from generate token
type JWTResponse struct {
	Token string
	Time  time.Time
}

// Credentials a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// User response from database
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
