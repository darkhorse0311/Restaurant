package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/reynld/carbtographer/pkg/yelp"

	"github.com/gorilla/mux"
)

// Locations returns local businees that names match restuarants in our db
func (s *Server) Locations(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	lat, err := strconv.ParseFloat(params["lat"], 64)
	if err != nil {
		log.Fatal(err)
	}
	lon, err := strconv.ParseFloat(params["lon"], 64)
	if err != nil {
		log.Fatal(err)
	}

	businesses := yelp.GetLocations(lat, lon)
	json.NewEncoder(w).Encode(businesses)
}
