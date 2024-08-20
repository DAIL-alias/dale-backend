package models

import (
	"time"
	
)

type Alias struct {
	AliasID     int       `gorm:"primaryKey;autoIncrement;column: alias_id;unique;not null" json:"aliasId"`
	UserID      int       `gorm:"column: user_id;unique;not null"                           json:"userId" `
	AliasPrefix string    `gorm:"column: alias_prefix;unique;not null"                      json:"aliasPrefix"`
	IsActive    bool      `gorm:"column: is_active;not null;default: true"                  json:"isActive"`
	IsDeleted   bool      `gorm:"column: is_deleted;not null;default:false"                 json:"isDeleted"`
	ExpiresAt   time.Time `gorm:"column: expires_at;not null"                               json:"expiresAt"`
	CreatedAt   time.Time `gorm:"column: created_at;not null"                               json:"createdAt"`
}
