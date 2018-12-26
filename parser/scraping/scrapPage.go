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

	return &database.AuchanProduct{
		URL:          url,
		Name:         mainProduct.Name,
		OldPrice:     mainProduct.OldPrice,
		CurrentPrice: mainProduct.Price,
		Quantity:     mainProduct.Quantity,
		ImageURL:     mainProduct.Gallery.Images[0].BigURL,
	}, nil
}

// GetProductCategory scrapes the product's category.
func GetProductCategory() {
	// TODO
}
