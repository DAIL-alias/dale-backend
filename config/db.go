package config

import (
	config "DALE/migrations"
	"DALE/models"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() {
	cstring := viper.GetString("C_STRING")
	if cstring == "" {
		log.Fatal("C_STRING not set")
	}

	// Start the connection
	var err error
	DB, err = gorm.Open(postgres.Open(cstring), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Migrate schema if needed
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Alias{})

	// Create triggers
	config.CreateTriggers(DB)
}