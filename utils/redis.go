package utils

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"PTH-IT/api_golang/config"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func RedisClient() *redis.Client {
	db, err := strconv.Atoi(config.Getconfig().Redis.Db)
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Getconfig().Redis.Host, config.Getconfig().Redis.Port),
		Password: config.Getconfig().Redis.Pass, // no password set
		Username: config.Getconfig().Redis.Username,
		DB:       db, // use default DB
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
		fmt.Println(err)
		return err
	}
	return nil
}
func DeleteToken(token string, userID string) bool {
	rdb := RedisClient()
	err := rdb.Del(ctx, userID).Err()
	if err != nil {
		return false
	}
	return true

}
