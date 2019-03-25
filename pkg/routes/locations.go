package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/machinebox/graphql"
	"github.com/reynld/carbtographer/pkg/models"
)

func searchBusiness(cl *graphql.Client, rest models.Restaurants, ch chan<- models.YelpResponse) {
	key := os.Getenv("YELP_API_KEY")
	var res models.YelpResponse

	req := graphql.NewRequest(yelpQuery)
	req.Var("name", rest.Name)
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

	fmt.Printf("\n%s\n%v\n", rest.Name, res)
	ch <- res
}

func getLocations(w http.ResponseWriter, req *http.Request) {
	var names []models.Restaurants
	db.Find(&names)
	// Length of names
	ln := len(names)

	client := graphql.NewClient("https://api.yelp.com/v3/graphql")
	// all businesses
	var ab []models.Business
	// unique id
	var uid []string

	c := make(chan models.YelpResponse)
	var wg sync.WaitGroup

	fmt.Printf("LEN : %d\n", ln)

	for i, name := range names {
		wg.Add(1)
		go func(n models.Restaurants, m int) {
			time.Sleep(time.Millisecond * time.Duration(150*m))
			searchBusiness(client, n, c)
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
			wg.Done()
		}
	}()

	wg.Wait()
	json.NewEncoder(w).Encode(&ab)
}
