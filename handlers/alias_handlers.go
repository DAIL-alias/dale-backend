package handlers

import (
	"DALE/config"
	"DALE/models"
	"DALE/services"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AliasHandler struct {
	AliasService *services.AliasService
}

func NewAliasHandler(aliasService *services.AliasService) *AliasHandler {
	return &AliasHandler{AliasService: aliasService}
}

func (h *AliasHandler) CreateAlias(c *gin.Context) {
	var alias models.Alias

	if err := c.ShouldBindJSON(&alias); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AliasService.CreateAlias(&alias); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alias)

}

func (h *AliasHandler) GetAliases(c *gin.Context) {
	users, err := h.AliasService.GetAliases()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *AliasHandler) GetAliasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	alias, err := h.AliasService.GetAliasByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Alias not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, alias)
}

func (h *AliasHandler) GetUsersAliases(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	aliases, err := h.AliasService.AliasRepository.GetUsersAliases(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aliases)
}

// Non-admin, enforces user resource ownership
// regular one gets all aliases, i moved those to an admin protected route
func (h *AliasHandler) GetUsersAliasesProtected(c *gin.Context) {
	// Get sessionID => userID
	sessionID, err := c.Cookie("sid")
	if err != nil {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		return
	}

	// Get userID from session
	userID, err := config.RedisClient.Get(context.Background(), sessionID).Result()
	if err == redis.Nil || userID == "" {
		// Invalid session
		c.Redirect(http.StatusFound, "/signin")
		c.Abort()
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// To integer
	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	aliases, err := h.AliasService.AliasRepository.GetUsersAliases(userIDint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aliases)
}

// define activation and deactivation of aliases
func (h *AliasHandler) ToggleActivateStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	
	// Get sessionID => userID
	sessionID, err := c.Cookie("sid")
	if err != nil {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		return
	}

	// Get userID from session
	userIDStr, err := config.RedisClient.Get(context.Background(), sessionID).Result()
	if err == redis.Nil || userIDStr == "" {
		// Invalid session
		c.Redirect(http.StatusFound, "/signin")
		c.Abort()
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	//fetch alias by id
	alias, err := h.AliasService.GetAliasByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if the userid on the alias is equal to the given id
	if alias.UserID != userID {
		// if not we return a invalid credentials error then continue with life
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	alias, err = h.AliasService.ToggleActiveStatus(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Alias not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, alias)
}
