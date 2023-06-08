package bootstrap

import (
	"Gin_Start/global"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func InitializeRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: global.App.Config.Redis.Host + ":" + global.App.Config.Redis.Port,
		//Addr:     "127.0.0.1:6379",
		Password: global.App.Config.Redis.Password,
		DB:       global.App.Config.Redis.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.Log.Error("Redis connect ping failed,err :", zap.Any("err", err))
		return nil, err
	}
	return client, nil
}
