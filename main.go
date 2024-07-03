package main

import (
	"github.com/gin-gonic/gin"
	"simple-go-gin-rest-api/infrastructure"
	"simple-go-gin-rest-api/routes"
)

func main() {
	infrastructure.Initialize()

	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
