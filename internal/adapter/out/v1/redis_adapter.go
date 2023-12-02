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

func (r *RedisAdapter) Count(ctx context.Context, businessId string, messageId string) (int, error) {
	val, err := r.rdb.Get(ctx, fmt.Sprintf("likit:%s:%s:count", businessId, messageId)).Int()
	if err == redis.Nil {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}
	return val, nil
}

func (r *RedisAdapter) Vote(ctx context.Context, businessId string, messageId string, userId string) error {
	keyForCount := fmt.Sprintf("likit:%s:%s:count", businessId, messageId)
	keyForVotedUser := fmt.Sprintf("likit:%s:%s:voted_user", businessId, messageId)

	sAdd := r.rdb.SAdd(ctx, keyForVotedUser, userId)

	if _, err := sAdd.Result(); err != nil {
		return err
	}

	if value, _ := sAdd.Result(); value == 1 {
		incr := r.rdb.Incr(ctx, keyForCount)
		if _, err := incr.Result(); err != nil {
			return err
		}
	}

	return nil
}

func (r *RedisAdapter) UnVote(ctx context.Context, businessId string, messageId string, userId string) error {
	keyForCount := fmt.Sprintf("likit:%s:%s:count", businessId, messageId)
	keyForVotedUser := fmt.Sprintf("likit:%s:%s:voted_user", businessId, messageId)

	sRem := r.rdb.SRem(ctx, keyForVotedUser, userId)

	if _, err := sRem.Result(); err != nil {
		return err
	}

	if value, _ := sRem.Result(); value == 1 {
		decr := r.rdb.Decr(ctx, keyForCount)
		if _, err := decr.Result(); err != nil {
			return err
		}
	}

	return nil
}

func (r *RedisAdapter) VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error) {
	members, err := r.rdb.SMembers(ctx, fmt.Sprintf("likit:%s:%s:voted_user", businessId, messageId)).Result()
	return members, err
}

func (r *RedisAdapter) IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error) {
	val, err := r.rdb.SIsMember(ctx, fmt.Sprintf("likit:%s:%s:voted_user", businessId, messageId), userId).Result()
	return val, err
}
