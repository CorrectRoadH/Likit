package server

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/in"
)

type VoteServer struct {
	businessIdMapVoteSystem map[string]in.VoteUseCase
	adminUseCase            in.AdminUseCase
}

func NewVoteServer(adminUseCase in.AdminUseCase) (in.VoteUseCase, error) {
	businessIdMapVoteSystem := make(map[string]in.VoteUseCase)

	// loading business data
	businesses, err := adminUseCase.Businesses(context.Background())
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
		adminUseCase:            adminUseCase,
	}, nil
}

func (v *VoteServer) Vote(ctx context.Context, businessId string, messageId string, userId string) (int64, error) {
	voteSystem, ok := v.businessIdMapVoteSystem[businessId]
	if !ok {
		return 0, domain.ErrBusinessNotExist
	}
	return voteSystem.Vote(ctx, businessId, messageId, userId)
}

func (v *VoteServer) UnVote(ctx context.Context, businessId string, messageId string, userId string) (int64, error) {
	voteSystem, ok := v.businessIdMapVoteSystem[businessId]
	if !ok {
		return 0, domain.ErrBusinessNotExist
	}
	return voteSystem.Vote(ctx, businessId, messageId, userId)
}

func (v *VoteServer) Count(ctx context.Context, businessId string, messageId string) (int64, error) {
	voteSystem, ok := v.businessIdMapVoteSystem[businessId]
	if !ok {
		return 0, domain.ErrBusinessNotExist
	}
	return voteSystem.Count(ctx, businessId, messageId)
}

func (v *VoteServer) IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error) {
	voteSystem, ok := v.businessIdMapVoteSystem[businessId]
	if !ok {
		return false, domain.ErrBusinessNotExist
	}
	return voteSystem.IsVoted(ctx, businessId, messageId, userId)
}

func (v *VoteServer) VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error) {
	voteSystem, ok := v.businessIdMapVoteSystem[businessId]
	if !ok {
		return nil, domain.ErrBusinessNotExist
	}
	return voteSystem.VotedUsers(ctx, businessId, messageId)
}
