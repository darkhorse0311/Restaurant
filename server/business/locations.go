package business

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/reynld/carbtographer/server/models"
)

// GetLocations returns local businees that names match restuarants in our db
func GetLocations(db *sql.DB, cache *redis.Client, w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	lon := params["lon"]
	lat := params["lat"]
	key := fmt.Sprintf("%s:%s", lon, lat)

	value, err := cache.Get(key).Result()
	if err != redis.Nil {
		jsonRes := make([]models.Business, 0)
		json.Unmarshal([]byte(value), &jsonRes)
		json.NewEncoder(w).Encode(jsonRes)
		return
	}

	businesses, err := getLocations(db, lat, lon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cacheBusinesses, err := json.Marshal(businesses)
	if err != redis.Nil {
		log.Println(err)
	}

	value, err = cache.Set(key, string(cacheBusinesses), time.Second*time.Duration(172800)).Result()
	if err != redis.Nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(businesses)
}

// searchBusiness gets ran in goroutine and returns the response to channel when done
func searchBusiness(name string, ch chan<- models.FusionResponse, lat string, lon string) {
	key := os.Getenv("YELP_API")
	client := http.Client{}

	query := fmt.Sprintf(
		"https://api.yelp.com/v3/businesses/search?latitude=%s&longitude=%s&term=%s&radius=%d",
		lat,
		lon,
		name,
		500,
	)

	// log.Println(query)

	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Authorization", "Bearer "+key)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Status)
	var result models.FusionResponse
	json.NewDecoder(resp.Body).Decode(&result)

	ch <- result
}

// getLocations returns local businees that names match restuarants in our db
func getLocations(db *sql.DB, lat string, lon string) ([]models.Business, error) {

	names, err := models.GetNames(db)
	if err != nil {
		return nil, err
	}

	c := make(chan models.FusionResponse)
	var wg sync.WaitGroup

	for i, name := range names {
		wg.Add(1)
		go func(n models.Restaurants, m int) {
			time.Sleep(time.Millisecond * time.Duration(200*m))
			defer wg.Done()
			re := regexp.MustCompile(`\b \b`)
			urlName := re.ReplaceAllLiteralString(n.Name, "%20")
			searchBusiness(urlName, c, lat, lon)
		}(name, i)
	}

	var ab []models.Business // all businesses
	var uid []string         // unique id
	go func() {
		for yr := range c {
			for _, business := range yr.Businesses {
				exist := false
				for _, id := range uid {
					if id == business.ID {
						exist = true
					}
				}

				vn := false //valid name
				for _, name := range names {
					if name.Name == business.Name {
						vn = true
					}
				}

				if exist == false && vn == true {
					ab = append(ab, business)
					uid = append(uid, business.ID)
				}
			}
		}
	}()

	wg.Wait()

	for i, bus := range ab {
		id, err := models.GetRestaurantID(db, bus.Name)
		if err != nil {
			return nil, err
		}
		ab[i].RID = id
	}

	return ab, nil
}
