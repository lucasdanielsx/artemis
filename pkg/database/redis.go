package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var ctx = context.Background()

func RedisHealth() (string, error) {
	return connection().Ping(ctx).Result()
}

func SetValue(key string, value string) (string, error) {
	return connection().Set(ctx, key, value, 0).Result()
}

func GetValue(key string) (string, error) {
	return connection().Get(ctx, key).Result()
}

func connection() *redis.Client {
	uri, _ := viper.Get("REDIS_URI").(string)
	pwd, _ := viper.Get("REDIS_PWD").(string)

	return redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: pwd,
		DB:       0,
	})
}
