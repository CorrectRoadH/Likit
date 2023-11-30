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
	val, err := r.rdb.Get(ctx, fmt.Sprintf("%s:%s:count", businessId, messageId)).Int()
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (r *RedisAdapter) Vote(ctx context.Context, businessId string, messageId string, userId string) error {
	keyForCount := fmt.Sprintf("%s:%s:count", businessId, messageId)
	keyForVotedUser := fmt.Sprintf("%s:%s:voted_user", businessId, messageId)

	// 开始一个 Redis 事务
	pipe := r.rdb.TxPipeline()

	// 加入操作到事务
	incr := pipe.Incr(ctx, keyForCount)
	sAdd := pipe.SAdd(ctx, keyForVotedUser, userId)

	// 执行事务
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}

	// 检查每个操作的结果
	if _, err := incr.Result(); err != nil {
		return err
	}
	if _, err := sAdd.Result(); err != nil {
		return err
	}

	return nil
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
	members, err := r.rdb.SMembers(ctx, fmt.Sprintf("%s:%s:voted_user", businessId, messageId)).Result()
	return members, err
}

func (r *RedisAdapter) IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error) {
	val, err := r.rdb.SIsMember(ctx, fmt.Sprintf("%s:%s:voted_user", businessId, messageId), userId).Result()
	return val, err
}
