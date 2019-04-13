package yelp

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/reynld/carbtographer/pkg/database"

	"github.com/machinebox/graphql"
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
func searchBusiness(cl *graphql.Client, rest models.Restaurants, ch chan<- Response, lat *float64, lon *float64) {
	key := os.Getenv("YELP_API_KEY")

	var res Response

	req := graphql.NewRequest(YelpQuery)
	req.Var("name", rest.Name)
	req.Var("lat", lat)
	req.Var("lon", lon)
	req.Header.Add("Authorization", "Bearer "+key)
	ctx := context.Background()

	err := cl.Run(ctx, req, &res)
	if err != nil {
		log.Fatal(err)
	}

	ch <- res
}

// GetLocations returns local businees that names match restuarants in our db
func GetLocations(lat float64, lon float64) []models.Business {

	if lon == float64(-74.0060) && lat == float64(40.7128) {
		jsonRes := make([]models.Business, 0)
		json.Unmarshal(DefaultLocation, &jsonRes)
		return jsonRes
	}

	var names []models.Restaurants
	database.GetNames(&names)

	client := graphql.NewClient("https://api.yelp.com/v3/graphql")
	var ab []models.Business // all businesses
	var uid []string         // unique id

	c := make(chan Response)
	var wg sync.WaitGroup

	for i, name := range names {
		wg.Add(1)
		go func(n models.Restaurants, m int) {
			time.Sleep(time.Millisecond * time.Duration(150*m))
			defer wg.Done()
			searchBusiness(client, n, c, &lat, &lon)
		}(name, i)
	}

	go func() {
		for yr := range c {
			for _, business := range yr.Search.Business {
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
		var dbResp []models.Restaurants
		database.GetIDByName(bus.Name, &dbResp)
		ab[i].RID = dbResp[0].ID
	}

	return ab
}
