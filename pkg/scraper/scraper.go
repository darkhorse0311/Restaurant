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

var remove = []string{
	"Culver",
	"Krystal+Burger",
	"Qdoba",
	"Tijuana+Flats+Burrito+Builder",
	"YaYa",
	"Zaxby",
}

var rename = map[string]string{
	"KentuckyFriedChicken": "KFC",
	"In-N-Out":             "In-N-Out Burger",
	"Hardee's/CarlsJr.":    "CarlsJr",
}

var logos = map[string]string{
	"Arby's":          "https://i.imgur.com/ZZtDD7M.png",
	"Boston Market":   "https://i.imgur.com/2OfB4Xm.png",
	"Burger King":     "https://i.imgur.com/AMwVZFX.png",
	"Chick Fil A":     "https://i.imgur.com/ZKGJde3.png",
	"Chipotle":        "https://i.imgur.com/1eKcZVo.png",
	"Dairy Queen":     "https://i.imgur.com/DTMvyUb.png",
	"Del Taco":        "https://i.imgur.com/B8oRmq5.png",
	"Domino's":        "https://i.imgur.com/5VZg5sN.png",
	"El Pollo Loco":   "https://i.imgur.com/FSd4mrl.png",
	"Firehouse Subs":  "https://i.imgur.com/tKUZVdN.png",
	"Five Guys":       "https://i.imgur.com/jcjy1yW.png",
	"Carl's Jr":       "https://i.imgur.com/KtrWf5D.png",
	"In-N-Out Burger": "https://i.imgur.com/hHIwrbN.png",
	"Jack in the Box": "https://i.imgur.com/wuOZWUf.png",
	"Jersey Mike's":   "https://i.imgur.com/y7ZB7yC.png",
	"Jimmy John's":    "https://i.imgur.com/KRluZGT.png",
	"KFC":             "https://i.imgur.com/cjRw1pY.png",
	"McDonald's":      "https://i.imgur.com/ZcANVPe.png",
	"Moe's":           "https://i.imgur.com/49e5Bql.png",
	"Panda Express":   "https://i.imgur.com/OiETaHQ.png",
	"Popeyes Chicken": "https://i.imgur.com/ryYOY3d.png",
	"Potbelly":        "https://i.imgur.com/FWgwDAJ.png",
	"Quiznos":         "https://i.imgur.com/zdi03On.png",
	"Sonic Drive-In":  "https://i.imgur.com/kDzvEDn.png",
	"Starbucks":       "https://i.imgur.com/5ZQysI2.png",
	"Subway":          "https://i.imgur.com/iPihkvD.png",
	"Taco Bell":       "https://i.imgur.com/gDFXzSu.png",
	"Wendy's":         "https://i.imgur.com/FJTkLbe.png",
	"Whataburger":     "https://i.imgur.com/nnIRV3y.png",
}

// Scraper struct
type Scraper struct {
	URL         string
	Links       []string
	Restaurants []models.JSONRestaurant
}

func isException(s string) error {
	for _, r := range remove {
		if strings.Contains(s, r) {
			return fmt.Errorf("is in exception array")
		}
	}
	return nil
}

// This will get called for each HTML element found
func (s *Scraper) processLink(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	href, exists := element.Children().First().Attr("href")
	if exists && strings.Contains(href, "food.asp") {
		if err := isException(href); err == nil {
			s.Links = append(s.Links, fmt.Sprintf("%s/%s", url, href))
		}
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

			if val, ok := rename[resName]; ok {
				s.Restaurants[i].Name = val
			} else {
				s.Restaurants[i].Name = resName
			}

		}

		rows := document.Find("tbody").Children()
		for k := range rows.Nodes {

			row := rows.Eq(k).Children()

			_, err := strconv.ParseFloat(row.Eq(3).First().Text(), 32)
			if err != nil {
				continue
			}

			item := models.JSONItem{}

			for l := range row.Nodes {
				switch l {
				case 1:
					item.Name = row.Eq(l).First().Text()
				case 2:
					item.Type = row.Eq(l).First().Text()
				case 3:
					item.Protein = getMacro(row.Eq(l))
				case 4:
					item.Fats = getMacro(row.Eq(l))
				case 5:
					item.Carbs = getMacro(row.Eq(l))
				case 6:
					item.Calories = getMacro(row.Eq(l))
				case 7:
					item.CalPerPro = getMacro(row.Eq(l))
				case 8:
					item.Sodium = getMacro(row.Eq(l))
				}
			}

			s.Restaurants[i].Items = append(s.Restaurants[i].Items, item)
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
	f, err := os.Create(filepath.Join(pwd, "restaurantData.json"))
	utils.Check(err)

	defer f.Close()

	jsonBytes := []byte(jsonString)
	_, err = f.Write(jsonBytes)
	utils.Check(err)

}
