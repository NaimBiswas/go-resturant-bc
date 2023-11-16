package database

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"restaurent-mng-bc/config"
	"time"
)

var RedisClient *redis.Client

func Redis() {
	appConfig, _ := config.Config()

	client := redis.NewClient(&redis.Options{
		Addr:       appConfig.RedisAddress,
		Password:   appConfig.RedisPassword,
		DB:         0,
		MaxRetries: 3,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	ping, err := client.Ping(ctx).Result()
	if err != nil {
		_ = client.Close()
		panic(err)
	}
	fmt.Println(ping)
	RedisClient = client
}
