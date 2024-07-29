package initialize

import (
	"context"
	"fmt"

	"github.com/hoangxt/go-ecommerce-ba-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
		PoolSize: 10,         // connection pool size
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization error", zap.Error(err))
	}

	fmt.Println("Redis is running")
	global.Rdb = rdb

	redisExample()
}

func redisExample() {
	// set key
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Println("set key error", zap.Error(err))
		return
	}
	// get key
	val, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("get key error", zap.Error(err))
		return
	}
	fmt.Println("score", val)

	global.Logger.Info("Value score is: ", zap.String("score", val))
}
