package database

import (
	"context"
	"fmt"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/out"
	"github.com/redis/go-redis/v9"
)

type RedisAdapter struct{}

func (r *RedisAdapter) TestConnect(config domain.DatabaseConnectConfig) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port), // Redis 服务器地址
		Password: config.Password,                                // 没有密码则留空
		DB:       0,                                              // 使用默认 DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	return err
}

func NewRedisAdapter() out.RedisPort {
	return &RedisAdapter{}
}
