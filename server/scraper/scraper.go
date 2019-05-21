package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/reynld/carbtographer/server"
	"github.com/reynld/carbtographer/server/models"
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
	"Hardee's/CarlsJr.":    "Carls Jr",
	"SonicDrive-In":        "Sonic Drive-In",
}

var logos = map[string]string{
	"Arby's":          "https://i.imgur.com/ZZtDD7M.png",
	"BostonMarket":    "https://i.imgur.com/2OfB4Xm.png",
	"BurgerKing":      "https://i.imgur.com/AMwVZFX.png",
	"Chick Fil A":     "https://i.imgur.com/ZKGJde3.png",
	"Chipotle":        "https://i.imgur.com/1eKcZVo.png",
	"DairyQueen":      "https://i.imgur.com/DTMvyUb.png",
	"DelTaco":         "https://i.imgur.com/B8oRmq5.png",
	"Domino's":        "https://i.imgur.com/5VZg5sN.png",
	"ElPolloLoco":     "https://i.imgur.com/FSd4mrl.png",
	"FirehouseSubs":   "https://i.imgur.com/tKUZVdN.png",
	"FiveGuys":        "https://i.imgur.com/jcjy1yW.png",
	"Carl's Jr":       "https://i.imgur.com/KtrWf5D.png",
	"In-N-Out Burger": "https://i.imgur.com/hHIwrbN.png",
	"JackintheBox":    "https://i.imgur.com/wuOZWUf.png",
	"JerseyMike's":    "https://i.imgur.com/y7ZB7yC.png",
	"JimmyJohn's":     "https://i.imgur.com/KRluZGT.png",
	"KFC":             "https://i.imgur.com/cjRw1pY.png",
	"McDonald's":      "https://i.imgur.com/ZcANVPe.png",
	"Moe's":           "https://i.imgur.com/49e5Bql.png",
	"PandaExpress":    "https://i.imgur.com/OiETaHQ.png",
	"PopeyesChicken":  "https://i.imgur.com/ryYOY3d.png",
	"Potbelly":        "https://i.imgur.com/FWgwDAJ.png",
	"Quiznos":         "https://i.imgur.com/zdi03On.png",
	"Sonic Drive-In":  "https://i.imgur.com/kDzvEDn.png",
	"Starbucks":       "https://i.imgur.com/5ZQysI2.png",
	"Subway":          "https://i.imgur.com/iPihkvD.png",
	"TacoBell":        "https://i.imgur.com/gDFXzSu.png",
	"Wendy's":         "https://i.imgur.com/FJTkLbe.png",
	"Whataburger":     "https://i.imgur.com/nnIRV3y.png",
}

// Scraper struct
type Scraper struct {
	URL         string
	Links       []string
	Restaurants []models.JSONRestaurant
}

func getMacro(sel *goquery.Selection) float32 {
	macro, err := strconv.ParseFloat(sel.First().Text(), 32)
	server.Check(err)
	return float32(macro)
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

func (s *Scraper) getInfo(link string, c chan<- models.Comu, index int) {
	response, err := http.Get(link)
	server.Check(err)

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	title := document.Find("head > title").First().Text()
	t := strings.Split(title, " ")

	resName := strings.Join(t[:len(t)-3], "")

	if val, ok := rename[resName]; ok {
		s.Restaurants[index].Name = val
		s.Restaurants[index].Logo = logos[val]
	} else {
		re := regexp.MustCompile(`\-`)
		newName := re.ReplaceAllLiteralString(resName, " ")
		fmt.Println(newName)
		s.Restaurants[index].Name = newName
		s.Restaurants[index].Logo = logos[newName]
	}

	rows := document.Find("tbody").Children()
	items := []models.JSONItem{}
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

		items = append(items, item)
	}
	c <- models.Comu{Items: items, Index: index}
}

func (s *Scraper) getItems() {
	c := make(chan models.Comu)
	var wg sync.WaitGroup

	for index, link := range s.Links {
		wg.Add(1)
		go func(l string, i int) {
			defer wg.Done()
			s.getInfo(l, c, i)
		}(link, index)
	}

	go func() {
		for co := range c {
			s.Restaurants[co.Index].Items = co.Items
		}
	}()

	wg.Wait()
}

// RunScraper get restaurant macro info
func RunScraper() {
	fmt.Println("Scraping starting...")
	s := Scraper{}
	response, err := http.Get(url)
	server.Check(err)

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	server.Check(err)
	document.Find(".pushy-submenu > ul > li").Each(s.processLink)
	fmt.Println("Retrieved all links...")

	res := make([]models.JSONRestaurant, len(s.Links))
	s.Restaurants = res
	s.getItems()
	fmt.Println("Scraped info successfully")

	jsonString, err := json.Marshal(s.Restaurants)
	server.Check(err)

	pwd, _ := os.Getwd()
	f, err := os.Create(filepath.Join(pwd, "restaurantData.json"))
	server.Check(err)
	fmt.Println("Writing file....")

	defer f.Close()

	jsonBytes := []byte(jsonString)
	_, err = f.Write(jsonBytes)
	server.Check(err)
	fmt.Println("Done!")

}
