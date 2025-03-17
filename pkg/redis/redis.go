package redis

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/oauth2"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func SaveTokenToRedis(c *gin.Context, token *oauth2.Token) error {
	data, err := json.Marshal(token)
	if err != nil {
		c.JSON(400, gin.H{"error": "can't save token to redis"})
		return fmt.Errorf("can't save token to redis: %w", err)
	}

	err = redisClient.Set(c, "token", data, 0).Err()
	if err != nil {
		c.JSON(400, gin.H{"error": "can't save token to redis"})
		return fmt.Errorf("can't save token to redis: %w", err)
	}

	return nil
}
