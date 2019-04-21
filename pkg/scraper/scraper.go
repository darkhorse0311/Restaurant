package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/reynld/carbtographer/pkg/models"
	"github.com/reynld/carbtographer/pkg/utils"
)

const url = "http://fastfoodmacros.com"

// Scraper struct
type Scraper struct {
	URL         string
	Links       []string
	Restaurants []models.JSONRestaurant
}

// This will get called for each HTML element found
func (s *Scraper) processLink(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	href, exists := element.Children().First().Attr("href")
	if exists && strings.Contains(href, "food.asp") {
		s.Links = append(s.Links, fmt.Sprintf("%s/%s", url, href))
	}
}

func (s *Scraper) getInfo() {

	for i, link := range s.Links {
		response, err := http.Get(link)
		utils.Check(err)

		document, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal("Error loading HTTP response body. ", err)
		}

		titles := document.Find("head > title")
		for j := range titles.Nodes {
			titleEl := titles.Eq(j).Text()
			t := strings.Split(titleEl, " ")

			resName := strings.Join(t[:len(t)-3], "")
			s.Restaurants[i].Name = resName
		}
		// fmt.Println(s.Restaurants[i].Name)

		rows := document.Find("tbody").Children()
		items := make([]models.JSONItem, len(rows.Nodes))
		s.Restaurants[i].Items = items

		for k := range rows.Nodes {

			fmt.Printf("K:%d\n", k)
			row := rows.Eq(k).Children()

			_, err := strconv.ParseFloat(row.Eq(3).First().Text(), 32)
			if err != nil {
				continue
			}

			for l := range row.Nodes {
				fmt.Printf("L:%d\n", l)

				switch l {
				case 1:
					s.Restaurants[i].Items[k].Name = row.Eq(l).First().Text()
				case 2:
					s.Restaurants[i].Items[k].Type = row.Eq(l).First().Text()
				case 3:
					s.Restaurants[i].Items[k].Protein = getMacro(row.Eq(l))
				case 4:
					s.Restaurants[i].Items[k].Fats = getMacro(row.Eq(l))
				case 5:
					s.Restaurants[i].Items[k].Carbs = getMacro(row.Eq(l))
				case 6:
					s.Restaurants[i].Items[k].Calories = getMacro(row.Eq(l))
				case 7:
					s.Restaurants[i].Items[k].CalPerPro = getMacro(row.Eq(l))
				case 8:
					s.Restaurants[i].Items[k].Sodium = getMacro(row.Eq(l))
				}

			}
		}
	}
}

func getMacro(sel *goquery.Selection) float32 {
	macro, err := strconv.ParseFloat(sel.First().Text(), 32)
	utils.Check(err)
	return float32(macro)
}

// RunScraper get restaurant macro info
func RunScraper() {
	s := Scraper{}
	response, err := http.Get(url)
	utils.Check(err)

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	utils.Check(err)

	document.Find(".pushy-submenu > ul > li").Each(s.processLink)

	res := make([]models.JSONRestaurant, len(s.Links))
	s.Restaurants = res
	s.getInfo()

	jsonString, err := json.Marshal(s.Restaurants)
	utils.Check(err)

	pwd, _ := os.Getwd()
	f, err := os.Create(filepath.Join(pwd, "testRestaurant.json"))
	utils.Check(err)

	defer f.Close()

	jsonBytes := []byte(jsonString)
	_, err = f.Write(jsonBytes)
	utils.Check(err)

}
