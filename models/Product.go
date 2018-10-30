package models

type Product struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Subtitle  string   `json:"subtitle"`
	Inventory string   `json:"inventory"`
	Options   []string `json:"options"`
	Price     string   `json:"price"`
	Image     []string `json:"image"`
}

type Products []Product
