package helpers

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

func ConfigRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func SetRedis(setname string, result []byte) error {
	redisClient := ConfigRedis()
	err := redisClient.Set(setname, result, 0).Err()
	if err != nil {
		return errors.New("failed to save results to redis")
	}

	err = redisClient.Expire("mykey", 7*24*time.Hour).Err()
	if err != nil {
		return err
	}

	time.Sleep(7 * 24 * time.Hour)

	return nil
}

func GetRedis(setname string) (string, bool, error) {
	redisClient := ConfigRedis()
	val, err := redisClient.Get(setname).Result()
	if err == redis.Nil {
		return "", true, errors.New("mykey does not exist")
	} else if err != nil {
		return "", false, err
	}

	return val, true, nil
}
