package handlers

import (
	"DALE/models"
	"DALE/services"
	"context"
	"net/http"
	"strconv"

	"log"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *services.AuthService
	UserService *services.UserService
}

func NewAuthHandler(authService *services.AuthService, userService *services.UserService) *AuthHandler {
	return &AuthHandler{AuthService: authService, UserService: userService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	// bind request json to user struct
	var loginReq struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate user using getuserbyemailandpassword service
	user, err := h.UserService.GetUserByEmailAndPassword(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
		return
	}
	userid := strconv.Itoa(int(user.ID))

	// create and return redis sesssion
	session, err := h.AuthService.CreateSession(context.Background(), userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set session ID as cookie for user
	c.SetCookie(
		"sid",
		session,
		services.SessionTTL,
		"/",
		"",
		false,
		true,
	)

	c.JSON(200, gin.H{
		"message": "Login successful",
	})
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	// get and bind user struct
	var newuser models.User

	log.Println(newuser)

	if err := c.ShouldBindJSON(&newuser); err != nil {
		// Bad request
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call create user service and return any errors
	if err := h.UserService.CreateUser(&newuser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// return created user
	c.JSON(200, gin.H{"message": "User created"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	session, err := c.Cookie("sid")
	if err != nil || session == "" {
		// Invalid session, 401 and redirect
		c.Redirect(http.StatusFound, "/signin")
		c.Abort()
		return
	}

	// delete the session if it exists
	err = h.AuthService.DeleteSession(context.Background(), session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session"})
		return
	}

	// Clear session cookie
	c.SetCookie("sid", "", -1, "/", "", false, true)

	//return success message with the deleted session information
	c.JSON(http.StatusOK, gin.H{
		"message": "Session successfully deleted",
	})
}
