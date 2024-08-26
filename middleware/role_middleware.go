package middleware

import (
	"DALE/config"
	"DALE/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RoleRequired(requiredRole int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get session ID
		session, err := c.Cookie("sid")
		if err != nil || session == "" {
			// Invalid session, 401 and redirect
			c.Redirect(http.StatusFound, "/signin")
			c.Abort()
			return
		}

		// Is SID valid?
		userID, err := config.RedisClient.Get(context.Background(), session).Result()
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

		// Fetch user role from database
		var user models.User
		db_err := config.DB.First(&user, userID).Error
		if db_err != nil {
			c.JSON(500, gin.H{"error": db_err.Error()})
			c.Abort()
			return
		}

		if user.Role < requiredRole {
			c.JSON(403, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
