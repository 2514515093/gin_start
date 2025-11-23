package utils

import (
	"gin_start/config"
	"time"
)

// Config Redis 配置
type RedisConfig struct {
	Addr     string // 127.0.0.1:6379
	Password string // 密码
	DB       int    // 0-15
}

// Set 设置值
func RdSetTime(key string, value interface{}, expire time.Duration) error {
	return config.Rdb.Set(config.Ctx, key, value, expire).Err()
}

func RdSet(key string, value interface{}) error {
	return config.Rdb.Set(config.Ctx, key, value, 2*time.Hour).Err()
}

// Get 获取值
func RdGet(key string) (string, error) {
	return config.Rdb.Get(config.Ctx, key).Result()
}
