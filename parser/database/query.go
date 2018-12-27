package database

import (
	"log"
	"fmt"

	"github.com/fatih/color"
	"github.com/jackc/pgx"
)

var (
	queryString = `INSERT INTO products ("product_id", "url", "name", "old_price", "current_price", "image_url", "category") VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT DO NOTHING`
)

// Query selects data from table and then sends record by record in channel.
func InsertIntoProducts(conn *pgx.ConnPool, product AuchanProduct) {
	_, err := conn.Exec(queryString, product.ProductID, product.URL, product.Name, product.OldPrice, product.CurrentPrice, product.ImageURL, product.Category)
	if err != nil {
		log.Println(err)
		return
	}
	output := fmt.Sprintf("PRODUCT %s IS SUCCESSFULLY ADDED...", product.ProductID)
	fmt.Println(color.GreenString(output))
}