package models

// AuchanProduct is a structure for keeping products in database.
type Product struct {
	ProductID    string   `json:"product_id"`
	URL          string   `json:"url"`
	Name         string   `json:"name"`
	OldPrice     int      `json:"old_price"`
	CurrentPrice int      `json:"current_price"`
	ImageURL     string   `json:"image_url"`
	Category     []string `json:"category"`
}