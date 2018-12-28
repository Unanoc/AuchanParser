package database

import (
	"server/errors"
	"server/models"

	"github.com/jackc/pgx"
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

func (db *DB) PostProductByIdHelper(product models.Product) (*models.Product, error) {

	rows := db.Conn.QueryRow(`
		INSERT
		INTO products ("product_id", "url", "name", "old_price", "current_price", "image_url", "category")
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING "product_id", "url", "name", "old_price", "current_price", "image_url", "category"`,
		product.ProductID,
		product.URL,
		product.Name,
		product.OldPrice,
		product.CurrentPrice,
		product.ImageURL,
		product.Category,
	)

	err := rows.Scan(
		&product.ProductID,
		&product.URL,
		&product.Name,
		&product.OldPrice,
		&product.CurrentPrice,
		&product.ImageURL,
		&product.Category,
	)
	if err != nil {
		switch err.(pgx.PgError).Code {
		case "23505":
			existProduct, _ := db.GetProductByIdHelper(product.ProductID)
			return existProduct, errors.ProdcutIsExist
		default:
			return nil, err
		}
	}

	return &product, nil
}

func (db *DB) GetProductsAllHelper(limit, priceFrom, priceTo string) (*models.Products, error) {
	products := models.Products{}
	var err error
	var rows *pgx.Rows

	if (priceFrom != "") && (priceTo != "") {
		rows, err = db.Conn.Query(`
		SELECT "product_id", "url", "name", "old_price", "current_price", "image_url", "category"
		FROM products
		WHERE current_price BETWEEN $1 AND $2
		ORDER BY "current_price"
		LIMIT $3`,
		priceFrom, priceTo, limit)
	} else {
		rows, err = db.Conn.Query(`
		SELECT "product_id", "url", "name", "old_price", "current_price", "image_url", "category"
		FROM products
		ORDER BY "current_price"
		LIMIT $1`,
		limit)
	}

	if err != nil {
		return nil, errors.ProductsNotFound
	}

	for rows.Next() {
		product := models.Product{}

		if err = rows.Scan(
			&product.ProductID,
			&product.URL,
			&product.Name,
			&product.OldPrice,
			&product.CurrentPrice,
			&product.ImageURL,
			&product.Category,
		); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return &products, nil
}

func (db *DB) GetProductsStatusHelper() (*models.Status, error) {
	status := models.Status{}

	row := db.Conn.QueryRow(`SELECT COUNT(*) FROM products`)
	err := row.Scan(&status.ProductsCount)
	if err != nil {
		return nil, err
	}

	return &status, nil
}