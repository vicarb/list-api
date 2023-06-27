package main

type Product struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	Stock       bool    `json:"stock" binding:"required"`
	Images      []Image `json:"images" binding:"required"`
}

type Image struct {
	URL string `json:"url" binding:"required"`
}
