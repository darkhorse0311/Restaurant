package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	"github.com/reynld/carbtographer/pkg/helpers"
	"github.com/reynld/carbtographer/pkg/routes"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := helpers.InitDB()

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// // Read
	// var product Product
	// db.First(&product, 1)                   // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212

	// // Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	// // Delete - delete product
	// db.Delete(&product)

	godotenv.Load()
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/", routes.GetServerIsUp).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(routes.RouteNotFound)
	fmt.Println("server live on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
