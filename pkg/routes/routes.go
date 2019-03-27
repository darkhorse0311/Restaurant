package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/reynld/carbtographer/pkg/database"

	"github.com/gorilla/mux"
	"github.com/reynld/carbtographer/pkg/models"
)

// GetItems returns all items per restuarant id
func GetItems(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var items []models.Items
	database.GetItems(params["id"], &items)
	json.NewEncoder(w).Encode(&items)
}

// GetNames returns all restaurant names in database
func GetNames(w http.ResponseWriter, req *http.Request) {
	var rest []models.Restaurants
	database.GetNames(&rest)
	json.NewEncoder(w).Encode(&rest)
}

// LoggingMiddleware logs HTTP request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
