package initialize

import (
	"context"

	"github.com/TaKieuLong/golang_fresher/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6381",
		Password: "",
		DB:       0,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization Error:", zap.Error(err))
	}
}
