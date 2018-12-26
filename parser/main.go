package main

import (
	"fmt"
	"parser/scraping"
)

func main() {
	a, err := scraping.GetProductInfo("https://www.auchan.ru/pokupki/antifriz-sintec-lux-g12-1-kg.html")
	if err == nil {
		fmt.Println(a)
	}
}
