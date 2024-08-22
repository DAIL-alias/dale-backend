package models

import (
	"DALE/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string `gorm:"column:email;unique;not null"                                 json:"email"     `
	Password   string `gorm:"column:password;not null"                                                      `
	Salt       string `gorm:"column:salt;not null"                                                          `
	NumAliases int    `gorm:"column: num_aliases;not null;default:0"                       json:"numAliases"`
}


// Hook to hash password and assign salt to the user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	salt, err := utils.GenerateSalt(32)

	// Check if there's an error
	if err != nil {
		return err
	}

	u.Password = utils.HashPassword(u.Password, salt)
	u.Salt = salt

	return nil
}
