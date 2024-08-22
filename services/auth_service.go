package services

// TTL for user session
const sessionTTL = 604800

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	redisClient *redis.Client
}

func NewAuthService(redisClient *redis.Client) *AuthService {
	return &AuthService{
		redisClient: redisClient
	}
}

// Create user session (login)
func (s *AuthService) CreateSession(ctx context.Context, userID string) error {
	return s.redisClient.Set(ctx, userID, 1, sessionTTL).Err()
}

// Delete user session
func (s *AuthService) DeleteSession(ctx context.Context, userID string) error {
	return s.redisClient.Del(ctx, userID).Err()
}

// Verify sid
func (s *AuthService) VerifySession(ctx context.Context, 