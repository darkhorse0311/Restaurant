package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reynld/carbtographer/pkg/database"
)

// GetItems returns all items per restuarant id
func (s *Server) GetItems(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	items, err := database.GetItems(s.DB, params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		return
	}

	json.NewEncoder(w).Encode(&items)
}

// GetNames returns all restaurant names in database
func (s *Server) GetNames(w http.ResponseWriter, req *http.Request) {
	rest, err := database.GetNames(s.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&rest)
}
