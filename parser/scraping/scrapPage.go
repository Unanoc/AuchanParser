package scraping

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"parser/database"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// GetProductInfo returns product info by URL.
func ProductInfoScrape(url string) (*database.AuchanProduct, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
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

	var productJSON ProductJSON
	err = json.Unmarshal(productJSONRaw, &productJSON)
	if err != nil {
		return nil, err
	}

	var mainProduct Product
	if product, ok := productJSON.Products[productJSON.MainProductID]; ok {
		mainProduct = product
	} else {
		return nil, fmt.Errorf("Cant find main product")
	}

	category, err := ProductCategoryScrape(url)
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
func ProductCategoryScrape(url string) ([]string, error) {
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

func ProductListScrape(url string) (*database.AuchanProduct, error) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	reg := regexp.MustCompile(`productListBlockJson\s*=\s*({.+});`)
	var productJsonRaw []byte
	match := reg.FindSubmatch(body)
	if len(match) > 1 {
		productJsonRaw = match[1]
	}
	var productListJson ProductListJSON

	err = json.Unmarshal(productJsonRaw, &productListJson)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v", productListJson)

	for _, product := range productListJson.Products {
		fmt.Println()
		fmt.Printf("URL: %v\n", product.URL)
		fmt.Printf("Name: %v\n", product.Name)
		fmt.Printf("Old Price: %v\n", product.OldPrice)
		fmt.Printf("New Price: %v\n", product.Price)
		fmt.Printf("Image: %v\n", product.Gallery.Images[0].NormalURL)
		fmt.Println()
	}

	fmt.Println("Current Page ", productListJson.ToolBarData.PagerData.CurrentPage)
	fmt.Println("Last Page ", productListJson.ToolBarData.PagerData.LastPage)
	return nil, nil
}