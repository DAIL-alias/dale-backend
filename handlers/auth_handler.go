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
		Email    string `json:"email" binding:"required"`
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
	var newuser models.User

	if err := c.ShouldBindJSON(&newuser); err != nil {
		// Bad request
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//call create user service and return any errors
	if err := h.UserService.CreateUser(&newuser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	//return created user
	c.JSON(200, newuser)
}

func (h *AuthHandler) LogoutUser(c *gin.Context) {
	//bind the request JSON to the req struct to fetch session ID and token
	var req struct {
		UserID       string `json:"user_id" binding:"required"`
		SessionToken string `json:"session_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Credentials"})
		return
	}

	//validate that the Redis session exists
	session, err := h.AuthService.VerifySession(context.Background(), req.SessionToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session token"})
		return
	}

	//delete the session if it exists
	err = h.AuthService.DeleteSession(context.Background(), req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session"})
		return
	}

	//return success message with the deleted session information
	c.JSON(http.StatusOK, gin.H{
		"message":         "Session successfully deleted",
		"deleted_session": session,
	})
}
