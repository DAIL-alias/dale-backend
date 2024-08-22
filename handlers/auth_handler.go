package handlers

import (
	_ "DALE/models"
	"DALE/services"

	"github.com/gin-gonic/gin"
)


type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

// Make endpoints for login, register, logout
// h is a reference to AuthHandler
// c is the context for the request (i.e. if you need params)
func (h *AuthHandler) Login(c *gin.Context) {
	//bind request json to user struct 
	

	//find user using email by apply getuserbyemail
	
	
	//comapre the password given and the hash password stored 

	//if hash of given == hash of stored then continue 
		// else return invalid credentials 
}
