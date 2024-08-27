package middleware

import (
	"DALE/config"
	"DALE/models"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RoleRequired(requiredRole int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get session ID
		session, err := c.Cookie("sid")
		log.Print(1, session, err)
		if err != nil || session == "" {
			// Invalid session, 401 and redirect
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Is SID valid?
		userID, err := config.RedisClient.Get(context.Background(), session).Result()
		log.Print(2, userID, err)
		if err == redis.Nil || userID == "" {
			// Invalid session
			c.JSON(401, gin.H{"error": "Unauthorized"})
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
		log.Print(3, user, db_err)
		if db_err != nil {
			c.JSON(500, gin.H{"error": db_err.Error(), "userID": userID})
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
