package server

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/CorrectRoadH/Likit/internal/port/out"
)

type VoteServer struct {
	voteStore out.SaveVoteUseCase
}

func NewVoteServer() in.VoteUseCase {
	return &VoteServer{}
}

func (v *VoteServer) Vote(ctx context.Context, businessId string, messageId string, userId string) error {
	return v.voteStore.Vote(ctx, businessId, messageId, userId)
}

func (v *VoteServer) Count(ctx context.Context, businessId, messageId, userId string) (int, error) {

	return v.voteStore.Count(ctx, businessId, messageId, userId)
}
