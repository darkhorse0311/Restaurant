package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reynld/carbtographer/pkg/models"
)

// getItems returns all items per restuarant
func getItems(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var items []models.Items
	// db.Find(&items, params["id"])
	db.Where("r_id = ?", params["id"]).Find(&items)
	json.NewEncoder(w).Encode(&items)
}
