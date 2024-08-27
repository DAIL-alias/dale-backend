package utils

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func UserIDFromSID(sid string, client *redis.Client) (int, error) {
	userID, err := client.Get(context.Background(), sid).Result()

	if err == redis.Nil || userID == "" {
		// Invalid session
		return -1, err
	} else if err != nil {
		// Unexpected
		return -1, err
	}

	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		// Unexpected
		return -1, err
	}

	return userIDint, nil
}
