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
	val, err := r.rdb.Get(ctx, fmt.Sprintf("%s:%s:count", businessId, messageId)).Int()
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (r *RedisAdapter) Vote(ctx context.Context, businessId string, messageId string, userId string) error {
	err := r.rdb.Incr(ctx, fmt.Sprintf("%s:%s:count", businessId, messageId))
	if err.Err() != nil {
		return err.Err()
	}
	err = r.rdb.SAdd(ctx, fmt.Sprintf("%s:%s:voted_user", businessId, messageId), userId)
	return err.Err()
}

func (r *RedisAdapter) UnVote(ctx context.Context, businessId string, messageId string, userId string) error {
	err := r.rdb.Decr(ctx, fmt.Sprintf("%s:%s:count", businessId, messageId))
	if err.Err() != nil {
		return err.Err()
	}
	err = r.rdb.SRem(ctx, fmt.Sprintf("%s:%s:voted_user", businessId, messageId), userId)
	return err.Err()
}

func (r *RedisAdapter) VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error) {
	members, err := r.rdb.SMembers(ctx, fmt.Sprintf("voted:user:%s:%s", businessId, messageId)).Result()
	return members, err
}

func (r *RedisAdapter) IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error) {
	val, err := r.rdb.SIsMember(ctx, fmt.Sprintf("%s:%s:voted_user", businessId, messageId), userId).Result()
	return val, err
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
