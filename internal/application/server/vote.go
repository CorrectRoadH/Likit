package server

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/in"
)

type VoteServer struct {
	businessIdMapVoteEngine map[string]in.VoteUseCase
	businessIdMapRankEngine map[string]in.RankVoteUseCase
	adminUseCase            in.AdminUseCase
}

func NewVoteServer(adminUseCase in.AdminUseCase) (in.VoteAdminUseCase, error) {
	businessIdMapVoteSystem := make(map[string]in.VoteUseCase)

	// loading business data
	businesses, err := adminUseCase.Businesses(context.Background())
	if err != nil {
		return nil, err
	}

	// create all vote system for business
	for _, business := range businesses {
		switch business.Type {
		case domain.SimpleVote:
			simpleVoteSystem, err := NewSimpleVoteSystem(business.Config)
			if err != nil {
				return nil, err
			}
			businessIdMapVoteSystem[business.Id] = simpleVoteSystem
		case domain.MiddleVote:

		default:
			return nil, domain.ErrVoteEngineNotSupport
		}
	}

	return &VoteServer{
		businessIdMapVoteEngine: businessIdMapVoteSystem,
		adminUseCase:            adminUseCase,
	}, nil
}

func (v *VoteServer) Vote(ctx context.Context, businessId string, messageId string, userId string) (int64, error) {
	voteSystem, ok := v.businessIdMapVoteEngine[businessId]
	if !ok {
		return 0, domain.ErrBusinessNotExist
	}
	return voteSystem.Vote(ctx, businessId, messageId, userId)
}

func (v *VoteServer) UnVote(ctx context.Context, businessId string, messageId string, userId string) (int64, error) {
	voteSystem, ok := v.businessIdMapVoteEngine[businessId]
	if !ok {
		return 0, domain.ErrBusinessNotExist
	}
	return voteSystem.Vote(ctx, businessId, messageId, userId)
}

func (v *VoteServer) Count(ctx context.Context, businessId string, messageId string) (int64, error) {
	voteSystem, ok := v.businessIdMapVoteEngine[businessId]
	if !ok {
		return 0, domain.ErrBusinessNotExist
	}
	return voteSystem.Count(ctx, businessId, messageId)
}

func (v *VoteServer) IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error) {
	voteSystem, ok := v.businessIdMapVoteEngine[businessId]
	if !ok {
		return false, domain.ErrBusinessNotExist
	}
	return voteSystem.IsVoted(ctx, businessId, messageId, userId)
}

func (v *VoteServer) VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error) {
	voteSystem, ok := v.businessIdMapVoteEngine[businessId]
	if !ok {
		return nil, domain.ErrBusinessNotExist
	}
	return voteSystem.VotedUsers(ctx, businessId, messageId)
}

func (v *VoteServer) CreateBusiness(ctx context.Context, business domain.Business) error {
	var simpleVoteEngine in.VoteUseCase
	var err error
	switch business.Type {
	case domain.SimpleVote:
		simpleVoteEngine, err = NewSimpleVoteSystem(business.Config)
	default:
		return domain.ErrVoteEngineNotSupport
	}

	if err != nil {
		return err
	}
	v.businessIdMapVoteEngine[business.Id] = simpleVoteEngine
	return nil
}
