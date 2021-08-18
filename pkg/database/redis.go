package database

import (
	"artemis/pkg/util"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var ctx = context.Background()

func RedisHealth() bool {
	_, err := redisConnection().Ping(ctx).Result()

	if err != nil {
		util.HandleError("Can not connect to redis: ", err)
		return false
	}

	return true
}

//func SetValue(key string, value string) (string, error) {
//	return redisConnection().Set(ctx, key, value, 0).Result()
//}
//
//func GetValue(key string) (string, error) {
//	return redisConnection().Get(ctx, key).Result()
//}

func redisConnection() *redis.Client {
	uri, _ := viper.Get("REDIS_URI").(string)
	pwd, _ := viper.Get("REDIS_PWD").(string)

	return redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: pwd,
		DB:       0,
	})
}
