package scraping

import (
	"sync"

	"github.com/jackc/pgx"
)

func Worker(wg *sync.WaitGroup, conn *pgx.ConnPool, urlsChan chan string) {
	defer wg.Done()

	for {
		select {
		case url, ok := <-urlsChan:
			if ok {
				go ProductPagesScrape(url, conn)
			} else {
				return
			}
		}
	}
}