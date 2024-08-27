package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func InitRedis() {
	addr := viper.GetString("REDIS_ADDR")
	pwd := viper.GetString("REDIS_PWD")

	if addr == "" || pwd == "" {
		log.Printf("Redis address: %s", addr)
		log.Printf("Redis password: %s", pwd)
		log.Fatal("redis config error")
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       0,
	})

	// Test connection
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Println("Redis connected")
}
