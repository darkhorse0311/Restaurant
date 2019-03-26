package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var r *mux.Router

// ConfigureRoutes sets routes for mux router
func ConfigureRoutes(database *gorm.DB, router *mux.Router) {
	r = router
	db = database

	r.HandleFunc("/", getServerIsUp).Methods("GET")

	r.HandleFunc("/names", getNames).Methods("GET")
	r.HandleFunc("/items/{id}", getItems).Methods("GET")
	r.HandleFunc("/locations/{lat}/{lon}", getLocations).Methods("GET")

	// lat: 40.7128, lon: -74.0060

	r.Use(loggingMiddleware)

	r.NotFoundHandler = http.HandlerFunc(routeNotFound)

}
