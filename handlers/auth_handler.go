package handlers

import (
	"DALE/models"
	_ "DALE/models"
	"DALE/services"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type AuthHandler struct {
	AuthService *services.AuthService
	UserService *services.UserService
}

func NewAuthHandler(authService *services.AuthService, userService *services.UserService) *AuthHandler {
	return &AuthHandler{AuthService: authService, UserService: userService}
}

// Make endpoints for login, register, logout
// h is a reference to AuthHandler
// c is the context for the request (i.e. if you need params)
func (h *AuthHandler) Login(c *gin.Context) {
	//bind request json to user struct 
	var loginReq struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validate user using getuserbyemailandpassword service
	user, err := h.UserService.GetUserByEmailAndPassword(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}
	userid := strconv.Itoa(int(user.ID))

	//create and return redis sesssion
	session, err := h.AuthService.CreateSession(context.Background(), userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, session)
}
 
func (h *AuthHandler) SignUp(c *gin.Context) {
	// get and bind user struct 
	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		// Bad request
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//call create user service and return any errors 
	// err := h.UserService.CreateUser(req)

	//return created user 
}