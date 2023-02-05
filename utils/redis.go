package utils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func RedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func GetToken(token string, userID string) bool {
	rdb := RedisClient()
	val, err := rdb.Get(ctx, userID).Result()
	if err == redis.Nil {
		return false
	}
	if val == token {
		return true
	}
	return false
}
func SetToken(token string, userID string) error {
	rdb := RedisClient()
	err := rdb.Set(ctx, userID, token, time.Duration(time.Minute*199)).Err()
	if err != nil {
		return err
	}
	return nil
}