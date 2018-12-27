package database

import (
	"server/errors"
	"server/models"
)

func (db *DB) GetProductByIdHelper(productID string) (*models.Product, error) {
	product := models.Product{}

	err := db.Conn.QueryRow(`
		SELECT "product_id", "url", "name", "old_price", "current_price", "image_url", "category"
		FROM products
		WHERE product_id = $1`,
		productID).Scan(
		&product.ProductID,
		&product.URL,
		&product.Name,
		&product.OldPrice,
		&product.CurrentPrice,
		&product.ImageURL,
		&product.Category,
	)

	if err != nil {
		return nil, errors.ProductNotFound
	}

	return &product, nil
}