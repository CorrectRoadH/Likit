package server

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/CorrectRoadH/Likit/internal/port/out"
)

type VoteServer struct {
	voteStore out.SaveVoteUseCase
}

func NewVoteServer(
	voteStore out.SaveVoteUseCase,
) in.VoteUseCase {
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
	panic("TODO: Implement")
}

func (v *VoteServer) UnVote(ctx context.Context, businessId string, messageId string, userId string) error {
	panic("TODO: Implement")
}

func (v *VoteServer) UserListOfVoted(ctx context.Context, businessId string, messageId string) ([]string, error) {
	panic("TODO: Implement")
}
