package database

// AuchanProduct is a structure for keeping products in database.
type AuchanProduct struct {
	URL          string
	Name         string
	OldPrice     int
	CurrentPrice int
	Quantity     int
	ImageURL     string
	// Category     []string
}
