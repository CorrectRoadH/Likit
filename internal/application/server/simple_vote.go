package server

import (
	"context"

	v1 "github.com/CorrectRoadH/Likit/internal/adapter/out/v1"
	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/CorrectRoadH/Likit/internal/port/out"
)

type SimpleVoteServer struct {
	voteStore out.SaveVoteUseCase
}

func NewSimpleVoteServer(voteStore out.SaveVoteUseCase) in.VoteUseCase {
	return &SimpleVoteServer{
		voteStore: voteStore,
	}
}

func (v *SimpleVoteServer) Vote(ctx context.Context, businessId string, messageId string, userId string) error {
	return v.voteStore.Vote(ctx, businessId, messageId, userId)
}

func (v *SimpleVoteServer) Count(ctx context.Context, businessId string, messageId string) (int, error) {
	return v.voteStore.Count(ctx, businessId, messageId)
}

func (v *SimpleVoteServer) IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error) {
	return v.voteStore.IsVoted(ctx, businessId, messageId, userId)
}

func (v *SimpleVoteServer) UnVote(ctx context.Context, businessId string, messageId string, userId string) error {
	return v.voteStore.UnVote(ctx, businessId, messageId, userId)
}

func (v *SimpleVoteServer) VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error) {
	return v.voteStore.VotedUsers(ctx, businessId, messageId)
}

func NewSimpleVoteSystem(config domain.Config) (in.VoteUseCase, error) {
	redisConfig, err := domain.GetRedisConfig(config)
	if err != nil {
		return nil, err
	}

	// check config require
	voteStore := v1.NewRedisAdapter(redisConfig)
	// s := NewSimpleVoteServer()
	return NewSimpleVoteServer(voteStore), nil
}
