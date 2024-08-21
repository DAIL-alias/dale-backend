package models

import (
	"time"

	"github.com/thanhpk/randstr"
	"gorm.io/gorm"
)

type Alias struct {
	gorm.Model
	UserID      int    `gorm:"column:user_id;not null"                                  json:"userID" `
	AliasPrefix string `gorm:"column:alias_prefix;unique;not null"                      json:"aliasPrefix"`
	IsActive    bool   `gorm:"column:is_active;not null;default: true"                  json:"isActive"`
	IsDeleted   bool   `gorm:"column:is_deleted;not null;default:false"                 json:"isDeleted"`
}

func (a *Alias) BeforeCreate(tx *gorm.DB) (err error) {
	alias := generateRandomString(7)

	a.AliasPrefix = alias
	a.CreatedAt = time.Now()

	return nil
}

func generateRandomString(length int) string {
	return randstr.String(length)
}
