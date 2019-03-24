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

func _searchBusiness(c *graphql.Client, rest models.Restaurants) models.YelpResponse {
	key := os.Getenv("YELP_API_KEY")
	var res models.YelpResponse

	req := graphql.NewRequest(yelpQuery)
	req.Var("name", rest.Name)
	req.Header.Add("Authorization", "Bearer "+key)

	ctx := context.Background()

	err := c.Run(ctx, req, &res)
	if err != nil {
		log.Fatal(err)
	}

	for i := range res.Search.Business {
		id := &res.Search.Business[i]
		id.RID = rest.ID
	}

	return res
}

func _getLocations(w http.ResponseWriter, req *http.Request) {
	url := os.Getenv("YELP_URL")
	var names []models.Restaurants
	db.Find(&names)
	// create a client (safe to share across requests)
	client := graphql.NewClient(url)
	// all businesses
	var ab []models.Business
	// unique id
	var uid []string
	for _, name := range names {
		yr := _searchBusiness(client, name)
		for _, business := range yr.Search.Business {
			var exist bool
			exist = false
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

	json.NewEncoder(w).Encode(&ab)

}
