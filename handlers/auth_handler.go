package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func signUp (c *gin.Context) {
	// Get the email
	var body struct {
		Email string
		Password string
		Username string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not read body",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func login (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}