package models

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    int     `json:"price"`
	Discount int     `json:"discount"`
	Rating   float64 `json:"rating"`
	Image    string  `json:"image"`
	Stock    int     `json:"stock"`
}
