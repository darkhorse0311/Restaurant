package scraper

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const url = "http://fastfoodmacros.com#"

// This will get called for each HTML element found
func processElement(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	href, exists := element.Children().First().Attr("href")
	if exists && strings.Contains(href, "food.asp") {
		fmt.Println(href)
	}
}

// RunScraper get restaurant macro info
func RunScraper() {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find(".pushy-submenu > ul > li").Each(processElement)

}
