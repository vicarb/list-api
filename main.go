package main

import (
	"github.com/gin-gonic/gin"
)

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

// Let's create a slice of Products to act as our database
var products = []Product{
	{
		Name:        "Example Product",
		Description: "This is an example product",
		Price:       9.99,
		Quantity:    100,
		Stock:       true,
		Images: []Image{
			{URL: "http://example.com/image1.png"},
			{URL: "http://example.com/image2.png"},
		},
	},
}

func main() {
	router := gin.Default()

	router.GET("/products", func(c *gin.Context) {
		c.JSON(200, products)
	})

	router.POST("/products", func(c *gin.Context) {
		var product Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(400, gin.H{
				"message": "Product data is not valid",
				"error":   err.Error(),
			})
			return
		}
		products = append(products, product)
		c.JSON(200, product)
	})

	router.Run(":8081") // Changed port to 8081
}
