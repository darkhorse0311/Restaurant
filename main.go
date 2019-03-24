package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	"github.com/reynld/carbtographer/pkg/helpers"
	"github.com/reynld/carbtographer/pkg/models"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

// GetServerIsUp '/' endpoint cheks if server is up
func getServerIsUp(w http.ResponseWriter, req *http.Request) {
	port := os.Getenv("PORT")

	res, _ := json.Marshal(struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{Status: 200, Message: "server live on port: " + port})

	w.Write(res)
}

// RouteNotFound '/*' endpoint for undefined routes
func routeNotFound(w http.ResponseWriter, req *http.Request) {
	res, _ := json.Marshal(struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{Status: 404, Message: "route not found"})

	w.Write(res)
}

func getNames(w http.ResponseWriter, req *http.Request) {
	var rest []models.Restaurants
	db.Find(&rest)

	json.NewEncoder(w).Encode(&rest)
}

func getItems(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var items []models.Items
	// db.Find(&items, params["id"])
	db.Where("r_id = ?", params["id"]).Find(&items)
	json.NewEncoder(w).Encode(&items)
}

var db *gorm.DB
var err error
var port string

func main() {
	db, err = helpers.InitDB()

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}
	r := mux.NewRouter()
	r.HandleFunc("/", getServerIsUp).Methods("GET")

	r.HandleFunc("/names", getNames).Methods("GET")
	r.HandleFunc("/items/{id}", getItems).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(routeNotFound)
	fmt.Println("server live on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
