package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	router := gin.Default()

	// Initialize the MongoDB client
	client := InitClient()
	collection := client.Database("mydatabase").Collection("products")

	router.GET("/products", func(c *gin.Context) {
		GetProductsHandler(c, collection)
	})

	router.POST("/products", func(c *gin.Context) {
		PostProductHandler(c, collection)
	})

	router.Run(":8081")
}
