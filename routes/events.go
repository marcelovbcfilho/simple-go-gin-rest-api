package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-go-gin-rest-api/models/event"
	"strconv"
)

func events(c *gin.Context) {
	events, err := event.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"events": events})
}

func eventById(c *gin.Context) {
	event, _, err := validateEventById(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	event := new(event.Event)
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
}

func updateEvent(c *gin.Context) {
	_, id, err := validateEventById(c)
	if err != nil {
		return
	}

	e := new(event.Event)
	err = c.ShouldBindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e, err = event.UpdateById(id, e)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}

func deleteEvent(c *gin.Context) {
	_, id, err := validateEventById(c)
	if err != nil {
		return
	}

	err = event.DeleteEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func validateEventById(c *gin.Context) (*event.Event, int64, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, 0, err
	}

	event, err := event.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, id, err
	}

	return event, id, err
}
