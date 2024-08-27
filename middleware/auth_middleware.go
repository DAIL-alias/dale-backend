package middleware

import (
	"DALE/config"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get session ID
		session, err := c.Cookie("sid")
		if err != nil || session == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Is SID valid?
		userID, err := config.RedisClient.Get(context.Background(), session).Result()
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

		// Valid session, proceed
		c.Next()
	}
}
