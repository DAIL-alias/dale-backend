package handlers

import (
	"DALE/models"
	"DALE/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

// define activation and deactivation of aliases
func (h *AliasHandler) ToggleActivateStatus(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	alias, err := h.AliasService.ToggleActiveStatus(id)
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
