package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/reynld/carbtographer/pkg/models"
	"github.com/reynld/carbtographer/pkg/yelp"

	"github.com/gorilla/mux"
)

// Locations returns local businees that names match restuarants in our db
func (s *Server) Locations(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	lon := params["lon"]
	lat := params["lat"]
	key := fmt.Sprintf("%s:%s", lon, lat)

	getCmd := s.Cache.Get(key)
	value, err := getCmd.Result()
	if err != redis.Nil {
		jsonRes := make([]models.Business, 0)
		json.Unmarshal([]byte(value), &jsonRes)
		json.NewEncoder(w).Encode(jsonRes)
		return
	}

	businesses, err := yelp.GetLocations(s.DB, lat, lon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cacheBusinesses, err := json.Marshal(businesses)
	if err != redis.Nil {
		log.Println(err)
	}

	setCmd := s.Cache.Set(key, string(cacheBusinesses), -1)
	value, err = setCmd.Result()
	if err != redis.Nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(businesses)
}
