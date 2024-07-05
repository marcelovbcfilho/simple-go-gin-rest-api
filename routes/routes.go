package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	//Event
	server.GET("/events", events)
	server.GET("/events/:id", eventById)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	//User
	server.POST("/signup", signUpUser)
	server.POST("/login", loginUser)
}
