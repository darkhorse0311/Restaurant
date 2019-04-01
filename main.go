package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/reynld/carbtographer/pkg/database"
	"github.com/reynld/carbtographer/pkg/routes"
)

func main() {
	godotenv.Load()
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		panic("enviroment variable PORT is required")
	}

	r := mux.NewRouter()
	r.Use(routes.LoggingMiddleware)

	r.HandleFunc("/", routes.GetServerIsUp).Methods("GET")
	r.HandleFunc("/names", routes.GetNames).Methods("GET")
	r.HandleFunc("/items/{id}", routes.GetItems).Methods("GET")
	r.HandleFunc("/locations/{lat}/{lon}", routes.GetLocations).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(routes.RouteNotFound)

	fmt.Printf("server live on port: %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))
}
