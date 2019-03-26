package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/machinebox/graphql"
	"github.com/reynld/carbtographer/pkg/models"
)

func searchBusiness(cl *graphql.Client, rest models.Restaurants, ch chan<- models.YelpResponse, lat *float64, lon *float64) {
	key := os.Getenv("YELP_API_KEY")
	var res models.YelpResponse

	req := graphql.NewRequest(yelpQuery)
	req.Var("name", rest.Name)
	req.Var("lat", lat)
	req.Var("lon", lon)
	req.Header.Add("Authorization", "Bearer "+key)
	ctx := context.Background()

	err := cl.Run(ctx, req, &res)
	if err != nil {
		log.Fatal(err)
	}

	for i := range res.Search.Business {
		id := &res.Search.Business[i]
		id.RID = rest.ID
	}

	// fmt.Printf("\n%s: %d", rest.Name, res.Search.Total)
	ch <- res
}

func getLocations(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	lat, _ := strconv.ParseFloat(params["lat"], 64)
	lon, _ := strconv.ParseFloat(params["lon"], 64)

	var names []models.Restaurants
	db.Find(&names)
	// Length of names
	// ln := len(names)
	// fmt.Printf("LEN : %d\n", ln)

	client := graphql.NewClient("https://api.yelp.com/v3/graphql")
	var ab []models.Business // all businesses
	var uid []string         // unique id

	c := make(chan models.YelpResponse)
	var wg sync.WaitGroup

	for i, name := range names {
		wg.Add(1)
		go func(n models.Restaurants, m int) {
			time.Sleep(time.Millisecond * time.Duration(150*m))
			defer wg.Done()
			// fmt.Printf("Started-%d: %s\n", m, n.Name)
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
				//valid name
				vn := false
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

	// fmt.Printf("\n\n\n------- BEFORE WAIT AND ENCODE -------\n\n\n")
	wg.Wait()
	json.NewEncoder(w).Encode(&ab)
	// fmt.Printf("\n\n\n------- AFTER WAIT AND ENCODE -------\n\n\n")
}
