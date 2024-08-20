package models

import (
	"time"
	
)

type User struct {
	UserID      int       `gorm:"column: user_id;unique;not null"                           json:"userId"    `
	Email       string    `gorm:"column: email;unique;not null"                             json:"email"     `
	Password    string    `gorm:"column: password;not null"                                                  `
	Salt        string    `gorm:"column: salt;not null                                                       `
	UpdatedAt   time.Time `gorm:"column: created_at;not null"                               json:"createdAt" `
	CreatedAt   time.Time `gorm:"column: created_at;not null"                               json:"createdAt" `
	NumAliases  int       `gorm:"column: num_aliases;not null"                              json:"numAliases"`
}
