package config

import (
	"github.com/spf13/viper"
    _ "DALE/models"
)

func LoadConfig() {
    viper.SetConfigFile(".env")
    err := viper.ReadInConfig()
    if err != nil {
        panic(err)
    }
}
