package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
	Ctx = context.Background()
)

func InitRedis() {
	// 创建客户端
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "211.159.169.85:6379", // redis地址
		Password: "",                    // 没有密码就写空 ""
		DB:       0,                     // 默认数据库
	})
	// 测试连接
	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis初始化成功")

}
