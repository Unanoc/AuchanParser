package types

// CategoryData extends ProductJSON.
type CategoryData struct {
	ID   string `json:"id"`
	Name string `json:"string"`
}

// Image extends ProductJSON.
type Image struct {
	Alt       string `json:"alt"`
	BigURL    string `json:"big"`
	NormalURL string `json:"image"`
	ThumbURL  string `json:"thumb"`
	RealURL   string `json:"real"`
	Selected  bool   `json:"selected"`
}

// Gallery extends ProductJSON.
type Gallery struct {
	Images []Image `json:"images"`
}

// Product extends ProductJSON.
type Product struct {
	ID       string  `json:"id"`
	IDGima   string  `json:"idGima"`
	InStock  bool    `json:"in_stock"`
	Name     string  `json:"name"`
	Price    int     `json:"price"`
	OldPrice int     `json:"old_price"`
	Quantity int     `json:"qty"`
	Gallery  Gallery `json:"gallery"`
}

// ProductJSON is a structure for unmarshalling of .js file with product's information.
type ProductJSON struct {
	MainProductCategory CategoryData       `json:"mainProductCategoryData"`
	MainProductID       string             `json:"mainProductId"`
	Products            map[string]Product `json:"products"`
}

// post
