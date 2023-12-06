package server

import (
	"context"
	"fmt"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/redis/go-redis/v9"
)

type SimpleVoteServer struct {
	rdb *redis.Client
}

func NewSimpleVoteServer(config domain.RedisConfig) in.VoteUseCase {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       0, // use default DB
	})

	return &SimpleVoteServer{
		rdb: rdb,
	}

}

func (v *SimpleVoteServer) Vote(ctx context.Context, businessId string, messageId string, userId string) (int, error) {
	keyForCount := fmt.Sprintf("likit:%s:%s:count", businessId, messageId)
	keyForVotedUser := fmt.Sprintf("likit:%s:%s:voted_user", businessId, messageId)

	voteNum := 0

	sAdd := v.rdb.SAdd(ctx, keyForVotedUser, userId)

	if _, err := sAdd.Result(); err != nil {
		return 0, err
	}

	if value, _ := sAdd.Result(); value == 1 {
		incr := v.rdb.Incr(ctx, keyForCount)
		if val, err := incr.Result(); err != nil {
			return 0, err
		} else {
			voteNum = int(val)
		}
	} else {
		return 0, domain.ErrUserHasVoted
	}
	// else the user has already voted

	return voteNum, nil
}

func (v *SimpleVoteServer) UnVote(ctx context.Context, businessId string, messageId string, userId string) (int, error) {
	keyForCount := fmt.Sprintf("likit:%s:%s:count", businessId, messageId)
	keyForVotedUser := fmt.Sprintf("likit:%s:%s:voted_user", businessId, messageId)

	voteNum := 0

	sRem := v.rdb.SRem(ctx, keyForVotedUser, userId)

	if _, err := sRem.Result(); err != nil {
		return 0, err
	}

	if value, _ := sRem.Result(); value == 1 {
		decr := v.rdb.Decr(ctx, keyForCount)
		if val, err := decr.Result(); err != nil {
			return 0, err
		} else {

			voteNum = int(val)
		}
	} else {
		return 0, domain.ErrUserNotVoted
	}
	// else the user not voted

	return voteNum, nil
}

func (v *SimpleVoteServer) Count(ctx context.Context, businessId string, messageId string) (int, error) {
	val, err := v.rdb.Get(ctx, fmt.Sprintf("likit:%s:%s:count", businessId, messageId)).Int()
	if err == redis.Nil {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}
	return val, nil
}

func (v *SimpleVoteServer) IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error) {
	val, err := v.rdb.SIsMember(ctx, fmt.Sprintf("likit:%s:%s:voted_user", businessId, messageId), userId).Result()
	return val, err
}

func (v *SimpleVoteServer) VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error) {
	members, err := v.rdb.SMembers(ctx, fmt.Sprintf("likit:%s:%s:voted_user", businessId, messageId)).Result()
	return members, err
}

func NewSimpleVoteSystem(config domain.Config) (in.VoteUseCase, error) {
	redisConfig, err := domain.GetRedisConfig(config)
	if err != nil {
		return nil, err
	}

	// s := NewSimpleVoteServer()
	return NewSimpleVoteServer(redisConfig), nil
}
