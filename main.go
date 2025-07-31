package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var items = []string{"Go", "Gin", "REST"}

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})
	r.GET("/items", getItems)
	r.POST("/items", addItem)
	log.Println("Server started on :8080")
	r.Run(":8080")
}

// GET /items
func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

// POST /items
func addItem(c *gin.Context) {
	var newItem struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items = append(items, newItem.Name)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Item added",
		"item":    newItem.Name,
	})
}
