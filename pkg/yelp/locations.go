package yelp

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

	"github.com/reynld/carbtographer/pkg/database"

	"github.com/reynld/carbtographer/pkg/models"
)

// Search db response
type Search struct {
	Total    int               `json:"total"`
	Business []models.Business `json:"business"`
}

// Response db response
type Response struct {
	Search Search `json:"search"`
}

// YelpQuery for business info
var YelpQuery = `query ($name: String!, $lat:Float, $lon: Float) {
	search(
		term: $name,
		latitude: $lat,
		longitude: $lon,
		radius: 500
	) {
		total
		business {
			id
			name
			coordinates {
				latitude
				longitude
			}
			photos
			distance
		}
	}
}`

// searchBusiness gets ran in goroutine and returns the response to channel when done
func searchBusiness(name string, ch chan<- FusionResponse, lat string, lon string) {
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
	var result FusionResponse
	json.NewDecoder(resp.Body).Decode(&result)

	ch <- result
}

// GetLocations returns local businees that names match restuarants in our db
func GetLocations(db *sql.DB, lat string, lon string) ([]Business, error) {

	names, err := database.GetNames(db)
	if err != nil {
		return nil, err
	}

	c := make(chan FusionResponse)
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

	var ab []Business // all businesses
	var uid []string  // unique id
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
		id, err := database.GetRestaurantID(db, bus.Name)
		if err != nil {
			return nil, err
		}
		ab[i].RID = id
	}

	return ab, nil
}
