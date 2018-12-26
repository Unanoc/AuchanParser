package scraping

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"parser/database"
	"parser/types"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// GetProductInfo returns product info by URL.
func GetProductInfo(url string) (*database.AuchanProduct, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	reg := regexp.MustCompile(`productBlockJson\s*=\s*({.+});`)
	var productJSONRaw []byte
	match := reg.FindSubmatch(body)
	if len(match) > 1 {
		productJSONRaw = match[1]
	}

	var productJSON types.ProductJSON
	err = json.Unmarshal(productJSONRaw, &productJSON)
	if err != nil {
		log.Fatal(err)
	}

	var mainProduct types.Product
	if product, ok := productJSON.Products[productJSON.MainProductID]; ok {
		mainProduct = product
	} else {
		return nil, fmt.Errorf("Cant find main product")
	}

	category, err := GetProductCategory(url)
	if err != nil {
		return nil, err
	}

	return &database.AuchanProduct{
		URL:          url,
		Name:         mainProduct.Name,
		OldPrice:     mainProduct.OldPrice,
		CurrentPrice: mainProduct.Price,
		Quantity:     mainProduct.Quantity,
		ImageURL:     mainProduct.Gallery.Images[0].BigURL,
		Category:     category,
	}, nil
}

// GetProductCategory scrapes the product's category.
func GetProductCategory(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	categories := doc.Find(`.breadcrumbs__list span[itemprop="name"]`).Map(func(_ int, s *goquery.Selection) string {
		return s.Text()
	})

	return categories[1:], nil
}
