package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-rest-api/infrastructure"
	"go-gin-rest-api/models"
	"net/http"
)

func main() {
	infrastructure.Initialize()

	server := gin.Default()

	server.GET("/events", func(c *gin.Context) {
		events, err := models.GetAllEvents()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"events": events})
	})

	server.POST("/events", func(c *gin.Context) {
		event := new(models.Event)
		err := c.ShouldBindJSON(&event)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = event.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"event": event})
	})

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
