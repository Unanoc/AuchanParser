package database

import (
	"log"

	"github.com/jackc/pgx"
)

var (
	queryString = `INSERT INTO products ("url", "name", "old_price", "current_price", "image_url", "category") VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING`
)

// Query selects data from table and then sends record by record in channel.
func InsertIntoProducts(conn *pgx.ConnPool, product AuchanProduct) {
	_, err := conn.Exec(queryString, product.URL, product.Name, product.OldPrice, product.CurrentPrice, product.ImageURL, product.Category)
	if err != nil {
		log.Println(err)
		return
	}
}