package cache

import (
	"context"
	"forum/logger"
	"log"

	"github.com/go-redis/redis/v8"
)

// RedisClient Redis缓存客户端单例
var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func Init() {
	// 创建上下文
	ctx := context.Background()
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 密码，如果没有则留空
		DB:       0,                // 默认数据库，默认为 0
	})
	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("无法连接到 Redis：%v", err)
	}
	RedisClient = rdb
	logger.Info("redis connect success")
}
