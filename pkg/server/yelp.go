package server

import (
	"encoding/json"
	"net/http"

	"github.com/reynld/carbtographer/pkg/models"
	"github.com/reynld/carbtographer/pkg/yelp"

	"github.com/gorilla/mux"
)

// Locations returns local businees that names match restuarants in our db
func (s *Server) Locations(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	lon := params["lon"]
	lat := params["lat"]

	if lon == "-74.0060" && lat == "40.7128" {
		jsonRes := make([]models.Business, 0)
		json.Unmarshal(yelp.DefaultLocation, &jsonRes)
		json.NewEncoder(w).Encode(jsonRes)
		return
	}

	businesses, err := yelp.GetLocations(s.DB, lat, lon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(businesses)
}
