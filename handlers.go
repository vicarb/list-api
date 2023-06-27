package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-driver/bson"
	"github.com/mongodb/mongo-driver/mongo"
	"log"
)

func GetProductsHandler(c *gin.Context, collection *mongo.Collection) {	// Define an array in which you can store the decoded documents
	var results []bson.M

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// Call cur.Next(ctx) to get the next document in the result. If the end of the result is reached, the cur.Next() returns false
	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		// Add the result to the array
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.Background())

	c.JSON(200, results)
}

func PostProductHandler(c *gin.Context, collection *mongo.Collection) {	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{
			"message": "Product data is not valid",
			"error":   err.Error(),
		})
		return
	}
	insertResult, err := collection.InsertOne(context.Background(), product)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, insertResult)
}
