package models

import (
	"gorm.io/gorm"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

type User struct {
	gorm.Model
	Email       string    `gorm:"column: email;unique;not null"                                json:"email"     `
	Password    string    `gorm:"column: password;not null"                                                     `
	Salt        string    `gorm:"column: salt;not null                                                          `
	// UpdatedAt   time.Time `gorm:"column: created_at"                                        json:"createdAt" `
	// CreatedAt   time.Time `gorm:"column: created_at"                                        json:"createdAt" `
	NumAliases  int       `gorm:"column: num_aliases;not null;default:0"                       json:"numAliases"`
}

// `generateSalt(length int)` and `hashPassword(password, salt string)`

// Hook to hash password and assign salt to the user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	salt, err := generateSalt(32)

	// Check if there's an error
	if err != nil {
		return err
	}

    u.Password = hashPassword(u.Password, salt)
	u.Salt = salt

	return nil
}

// Generates a random salt of a length
// May return an error, propgaga
func generateSalt(length int) (string, error) {
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
func hashPassword(password, salt string) string {
    // Generate and write into hash
    hash := sha256.New()
    hash.Write([]byte(password + salt))

    // Convert to base64 encoding
    return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}