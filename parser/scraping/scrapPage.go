package scraping

import (
	"encoding/json"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"parser/database"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/jackc/pgx"
)

func ProductPagesScrape(url string, conn *pgx.ConnPool) {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
	}

	for _, product := range productListJson.Products {
		category, err := ProductCategoryScrape(product.URL)
		if err != nil {
			log.Println(err)
		}
		auchanItem := database.AuchanProduct{
			ProductID: product.ID,
			URL: product.URL,
			Name: product.Name,
			OldPrice: product.OldPrice,
			CurrentPrice: product.Price,
			ImageURL: product.Gallery.Images[0].NormalURL,
			Category: category,
		}

		database.InsertIntoProducts(conn, auchanItem)
	}
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

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	categories := doc.Find(`.breadcrumbs__list span[itemprop="name"]`).Map(func(_ int, s *goquery.Selection) string {
		return s.Text()
	})

	return categories[1:], nil
}
