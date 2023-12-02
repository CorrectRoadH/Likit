package server

import (
	"context"
	"fmt"

	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/CorrectRoadH/Likit/internal/port/out"
)

type VoteServer struct {
	businessIdMapVoteSystem map[string]in.VoteUseCase
	businessStore           out.BusinessUseCase
}

func NewVoteServer(businessUseCase out.BusinessUseCase) (in.VoteUseCase, error) {
	businessIdMapVoteSystem := make(map[string]in.VoteUseCase)

	// loading business data
	businesses, err := businessUseCase.Businesses(context.Background())
	if err != nil {
		return nil, err
	}

	// create all vote system for business
	for _, business := range businesses {
		simpleVoteSystem, err := NewSimpleVoteSystem(business.Config)
		if err != nil {
			return nil, err
		}
		businessIdMapVoteSystem[business.Id] = simpleVoteSystem
	}

	return &VoteServer{
		businessIdMapVoteSystem: businessIdMapVoteSystem,
		businessStore:           businessUseCase,
	}, nil
}

func (v *VoteServer) Vote(ctx context.Context, businessId string, messageId string, userId string) error {
	voteSystem, ok := v.businessIdMapVoteSystem[businessId]
	if !ok {
		return fmt.Errorf("Business is not exist")
	}
	return voteSystem.Vote(ctx, businessId, messageId, userId)
}

func (v *VoteServer) UnVote(ctx context.Context, businessId string, messageId string, userId string) error {
	voteSystem, ok := v.businessIdMapVoteSystem[businessId]
	if !ok {
		return fmt.Errorf("Business is not exist")
	}
	return voteSystem.Vote(ctx, businessId, messageId, userId)
}

func (v *VoteServer) Count(ctx context.Context, businessId string, messageId string) (int, error) {
	voteSystem, ok := v.businessIdMapVoteSystem[businessId]
	if !ok {
		return 0, fmt.Errorf("Business is not exist")
	}
	return voteSystem.Count(ctx, businessId, messageId)
}

func (v *VoteServer) IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error) {
	voteSystem, ok := v.businessIdMapVoteSystem[businessId]
	if !ok {
		return false, fmt.Errorf("Business is not exist")
	}
	return voteSystem.IsVoted(ctx, businessId, messageId, userId)
}

func (v *VoteServer) VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error) {
	voteSystem, ok := v.businessIdMapVoteSystem[businessId]
	if !ok {
		return nil, fmt.Errorf("Business is not exist")
	}
	return voteSystem.VotedUsers(ctx, businessId, messageId)
}
