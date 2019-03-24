package routes

import (
	"encoding/json"
	"net/http"

	"github.com/reynld/carbtographer/pkg/models"
)

// getNames returns all restaurant names in database
func getNames(w http.ResponseWriter, req *http.Request) {
	var rest []models.Restaurants
	db.Find(&rest)

	json.NewEncoder(w).Encode(&rest)
}
