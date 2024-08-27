package utils

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func UserIDFromSID(sid string, client *redis.Client) (string, error) {
	userID, err := client.Get(context.Background(), sid).Result()

	if err == redis.Nil || userID == "" {
		// Invalid session
		return "", err
	} else if err != nil {
		// Unexpected
		return "", err
	}

	return userID, nil
}
