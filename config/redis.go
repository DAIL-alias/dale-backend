package config

import (
	"log"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func init() {
	addr := viper.GetString("REDIS_ADDR")
	pwd := viper.GetString("REDIS_PWD")

	if addr == "" || pwd == "" {
		log.Fatal("redis config error")
	}


	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
		Password: pwd,
		DB: 0,
	})
}