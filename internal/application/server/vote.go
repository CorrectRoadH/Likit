package server

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/CorrectRoadH/Likit/internal/port/out"
)

type VoteServer struct {
	voteStore out.SaveVoteUseCase
}

func NewVoteServer(voteStore out.SaveVoteUseCase) in.VoteUseCase {
	return &VoteServer{
		voteStore: voteStore,
	}
}

func (v *VoteServer) Vote(ctx context.Context, businessId string, messageId string, userId string) error {
	return v.voteStore.Vote(ctx, businessId, messageId, userId)
}

func (v *VoteServer) Count(ctx context.Context, businessId string, messageId string) (int, error) {

	return v.voteStore.Count(ctx, businessId, messageId)
}

func (v *VoteServer) IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error) {
	return v.voteStore.IsVoted(ctx, businessId, messageId, userId)
}

func (v *VoteServer) UnVote(ctx context.Context, businessId string, messageId string, userId string) error {
	return v.voteStore.UnVote(ctx, businessId, messageId, userId)
}

func (v *VoteServer) VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error) {
	return v.voteStore.VotedUsers(ctx, businessId, messageId)
}
