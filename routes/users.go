package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-go-gin-rest-api/models/user"
)

func signUpUser(c *gin.Context) {
	var user user.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func loginUser(c *gin.Context) {
}
