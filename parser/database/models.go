package database

// AuchanProduct is a structure for keeping products in database.
type AuchanProduct struct {
	ProductID    string
	URL          string
	Name         string
	OldPrice     int
	CurrentPrice int
	ImageURL     string
	Category     []string
}
