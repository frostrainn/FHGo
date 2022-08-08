package drivers

import (
	"context"
	"fhgo/global"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "zhangzhongyuan.cn:56379",
		Password: "zzy@3.14", // no password set
		DB:       0,          // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("Redis连接失败")
	} else {
		fmt.Println("Redis连接成功")
		global.Logger.Info("redis connect ping response:", zap.String("pong", pong))
		global.Redis = client
	}
}
