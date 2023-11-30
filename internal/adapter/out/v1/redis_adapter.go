package v1

import (
	"context"
	"fmt"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/out"
	"github.com/redis/go-redis/v9"
)

type RedisAdapter struct {
	rdb *redis.Client
}

func (r *RedisAdapter) Count(ctx context.Context, businessId string, messageId string) (int, error) {
	val, err := r.rdb.Get(ctx, fmt.Sprintf("count-%s-%s", businessId, messageId)).Int()
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (r *RedisAdapter) Vote(ctx context.Context, businessId string, messageId string, userId string) error {
	err := r.rdb.Incr(ctx, fmt.Sprintf("count-%s-%s", businessId, messageId))
	return err.Err()
}

func NewRedisAdapter(config domain.RedisConfig) out.SaveVoteUseCase {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Passwd,
		DB:       0, // use default DB
	})

	return &RedisAdapter{
		rdb: rdb,
	}
}