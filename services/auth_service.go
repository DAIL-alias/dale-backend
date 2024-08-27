package services

import (
	"DALE/repositories"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// TTL for user session
const SessionTTL = 0

type AuthService struct {
	redisClient    *redis.Client
	UserRepository *repositories.UserRepository
}

func NewAuthService(redisClient *redis.Client, userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{
		redisClient: redisClient,
	}
}

// Create user session (login)
func (s *AuthService) CreateSession(ctx context.Context, userID string) (string, error) {
	sessionToken := uuid.NewString()

	err := s.redisClient.Set(ctx, sessionToken, userID, SessionTTL).Err()
	if err != nil {
		return "", err
	}

	return sessionToken, nil
}

// Delete user session
func (s *AuthService) DeleteSession(ctx context.Context, sessionID string) error {
	err := s.redisClient.Del(ctx, sessionID).Err()

	if err != nil {
		return err
	}

	return nil
}

// Verify sid
func (s *AuthService) VerifySession(ctx context.Context, sessionToken string) (string, error) {
	// Get userID for session token
	userID, err := s.redisClient.Get(ctx, sessionToken).Result()

	if err == redis.Nil {
		// Implies session token invalid
		return "", errors.New("invalid token")
	} else if err != nil {
		// Unexpected
		return "", err
	}

	// To be handled later
	return userID, nil
}

// Get userID associated with sessionToken if any
func (s *AuthService) GetSession(ctx context.Context, sessionToken string) (string, error) {
	userID, err := s.redisClient.Get(ctx, sessionToken).Result()

	if err == redis.Nil {
		// Implies no session exists
		return "", nil
	} else if err != nil {
		// Unexpected
		return "", err
	}

	return userID, nil
}
