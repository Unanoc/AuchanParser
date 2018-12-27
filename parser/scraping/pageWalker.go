package scraping

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly"
)

func GetProductsURLs(urlsChan chan string) {
	c := colly.NewCollector(
		// Visit only root url and urls which start with "/pokupki/"
		colly.URLFilters(
			regexp.MustCompile(`https?://(www.)?auchan\.ru/pokupki/.+`),
			regexp.MustCompile(`https?://(www.)?auchan\.ru/$`),
		),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Visit link found on page
		// Only those links are visited which are matched by  any of the URLFilter regexps
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
		urlsChan <- r.URL.String()
	})

	c.Visit("https://www.auchan.ru/")
	fmt.Println("ЗДЕСЬ МОЖНО ЗАКРЫТЬ КАНАЛ!!!")
}