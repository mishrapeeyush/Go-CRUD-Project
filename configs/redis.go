package configs

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)
var (
    RedisClient *redis.Client
)
func InitRedis() {
    // Connect to Redis
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    // Test the connection
    _, err := RedisClient.Ping(context.Background()).Result()
    if err != nil {
        log.Fatal("Failed to connect to Redis:", err)
    }else{
		fmt.Printf("Connected to Redis")
	}
}