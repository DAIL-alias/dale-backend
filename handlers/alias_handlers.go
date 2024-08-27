package handlers

import (
	"DALE/config"
	"DALE/models"
	"DALE/services"
	"DALE/utils"
	"log"
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
	// Get session ID
	session, err := c.Cookie("sid")
	if err != nil || session == "" {
		// Invalid session, 401 and redirect
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Is SID valid?
	userID, err := utils.UserIDFromSID(session, config.RedisClient)
	if err == redis.Nil {
		// Invalid session
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var alias models.Alias
	// Create the alias model with UserID specified only
	alias.UserID = userID

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
	log.Println("GetUsersAliasesProtected")
	// Get sessionID => userID
	sessionID, err := c.Cookie("sid")
	if err != nil {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		return
	}

	// Get userID from session
	userID, err := utils.UserIDFromSID(sessionID, config.RedisClient)
	if err == redis.Nil {
		// Invalid session
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	aliases, err := h.AliasService.AliasRepository.GetUsersAliases(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aliases)
}

// define activation and deactivation of aliases
func (h *AliasHandler) ToggleActivateStatus(c *gin.Context) {
	aliasIDStr := c.Param("id")

	aliasID, err := strconv.Atoi(aliasIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	
	sessionID, err := c.Cookie("sid")
	if err != nil {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		return
	}

	// Get userID from session
	userID, err := utils.UserIDFromSID(sessionID, config.RedisClient)
	if err == redis.Nil {
		// Invalid session
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	//fetch alias by id
	alias, err := h.AliasService.GetAliasByID(aliasID)
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
	
	alias, err = h.AliasService.ToggleActiveStatus(aliasID)
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

// Delete alias
func (h *AliasHandler) DeleteAlias(c *gin.Context) {
	aliasIDStr := c.Param("id")

	aliasID, err := strconv.Atoi(aliasIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	
	sessionID, err := c.Cookie("sid")
	if err != nil {
		c.JSON(403, gin.H{"error": "Unauthorized"})
		return
	}

	// Get userID from session
	userID, err := utils.UserIDFromSID(sessionID, config.RedisClient)
	if err == redis.Nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	alias, err := h.AliasService.GetAliasByID(aliasID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if alias.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = h.AliasService.DeleteAlias(aliasID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alias deleted successfully"})
}