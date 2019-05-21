package business

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reynld/carbtographer/server/models"
)

// GetItems returns all items per restuarant id
func GetItems(db *sql.DB, w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	items, err := models.GetItems(db, params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		return
	}

	json.NewEncoder(w).Encode(&items)
}

// GetNames returns all restaurant names in database
func GetNames(db *sql.DB, w http.ResponseWriter, req *http.Request) {
	rest, err := models.GetNames(db)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&rest)
}
