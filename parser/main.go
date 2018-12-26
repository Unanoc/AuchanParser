package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// ProductPageScrape returns product info by url.
func ProductPageScrape(url string) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the name of product
	doc.Find(".prcard__title").Each(func(i int, s *goquery.Selection) {
		productName := s.Find("h1").Text()
		fmt.Printf("%s\n", productName)
	})

	// Finc the price of product
	doc.Find(".prcard-current-price").Each(func(i int, s *goquery.Selection) {
		productPriceStr := s.Find("span").Text()
		fmt.Printf("%s\n", productPriceStr)
	})

	// Finc the price of image
	doc.Find(".zoom_container .zoomImg img").Each(func(i int, s *goquery.Selection) {
		imageURL, ok := s.Attr("src")
		if ok {
			fmt.Printf("%s\n", imageURL)
		}
	})
}

func main() {
	ProductPageScrape("https://www.auchan.ru/pokupki/antifriz-sintec-lux-g12-1-kg.html")
}
