package handlers

import (
	"net/http"
	"DALE/models"
	"DALE/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserHandler to handle requests for user stuff
type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// Handle user creation
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	// Bind req body to user struct if possible
	if err := c.ShouldBindJSON(&user); err != nil {
		// Bad request
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	if err := h.UserService.CreateUser(&user); err != nil {
		// Some internal server error
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Return user
	c.JSON(http.StatusOK, user)
}

// Retrieve all users
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.UserService.GetUsers()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Retrieve user by ID
func (h *UserHandler) GetUserById(c *gin.Context) {
	id_str := c.Param("id")  // Get ID parameter from request
	
	// Convert ID to integer
	id, err := strconv.Atoi(id_str)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	
	user, err := h.UserService.GetUserById(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}