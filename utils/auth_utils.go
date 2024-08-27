package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

// Generates a random salt of a length
// May return an error, propgagate
func GenerateSalt(length int) (string, error) {
	// Allocate byte slice
	salt := make([]byte, length)

	// Read random bytes into slice
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}

	// Encode salt to base64
	return base64.StdEncoding.EncodeToString(salt), nil
}

// Hashes a password with a given salt
func HashPassword(password, salt string) string {
	// Generate and write into hash
	hash := sha256.New()
	hash.Write([]byte(password + salt))

	// Convert to base64 encoding
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// Checks if password matches other hash
func CheckPasswordHash(password, salt, otherHash string) bool {
	return HashPassword(password, salt) == otherHash
}