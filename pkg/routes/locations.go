package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/machinebox/graphql"
	"github.com/reynld/carbtographer/pkg/models"
)

func searchBusiness(cl *graphql.Client, s string, ch chan models.YelpResponse) {
	fmt.Printf("started: %s\n", s)
	key := os.Getenv("YELP_API_KEY")
	var res models.YelpResponse

	req := graphql.NewRequest(yelpQuery)
	req.Var("name", s)
	req.Header.Add("Authorization", "Bearer "+key)

	ctx := context.Background()

	err := cl.Run(ctx, req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ended: %s\n", s)
	ch <- res
}

func getLocations(w http.ResponseWriter, req *http.Request) {
	url := os.Getenv("YELP_URL")
	var names []models.Restaurants
	db.Find(&names)
	// create a client (safe to share across requests)
	client := graphql.NewClient(url)
	var ab []models.Business
	var us []string

	c := make(chan models.YelpResponse)

	for _, name := range names {
		go searchBusiness(client, name.Name, c)
	}

	count := 0
	for yr := range c {
		// ab = append(ab, yr.Search.Business...)
		for _, business := range yr.Search.Business {
			var exist bool
			exist = false
			for _, id := range us {
				if id == business.ID {
					exist = true
				}
			}
			if exist == false {
				ab = append(ab, business)
				us = append(us, business.ID)
			}
		}
		count++
		if count >= len(names) {
			close(c)
		}
	}

	json.NewEncoder(w).Encode(&ab)

}
