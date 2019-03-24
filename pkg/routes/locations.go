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

func searchBusiness(c *graphql.Client, s string) models.YelpResponse {
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

func getLocations(w http.ResponseWriter, req *http.Request) {
	url := os.Getenv("YELP_URL")
	// var names []models.Restaurants
	// db.Find(&names)

	// create a client (safe to share across requests)
	client := graphql.NewClient(url)
	yr := searchBusiness(client, "Mcdonalds")

	json.NewEncoder(w).Encode(&yr)

	// }

	// set headers

	// set any variables

	// run and capture the response
	// var respData ResponseStruct
	// if err := client.Run(ctx, r, &respData); err != nil {
	// 	log.Fatal(err)
	// }

}
