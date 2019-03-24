package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/machinebox/graphql"
	"github.com/reynld/carbtographer/pkg/models"
)

func _searchBusiness(c *graphql.Client, s string) models.YelpResponse {
	key := os.Getenv("YELP_API_KEY")
	var res models.YelpResponse

	req := graphql.NewRequest(yelpQuery)
	req.Var("name", s)
	req.Header.Add("Authorization", "Bearer "+key)

	ctx := context.Background()

	err := c.Run(ctx, req, &res)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func _getLocations(w http.ResponseWriter, req *http.Request) {
	url := os.Getenv("YELP_URL")
	var names []models.Restaurants
	db.Find(&names)
	// create a client (safe to share across requests)
	client := graphql.NewClient(url)
	var ab []models.Business
	var us []string

	for _, name := range names {
		yr := _searchBusiness(client, name.Name)
		for _, business := range yr.Search.Business {
			var exist bool
			exist = false
			for _, id := range us {
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
				us = append(us, business.ID)
			}
		}
	}

	json.NewEncoder(w).Encode(&ab)

}
