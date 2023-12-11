package in

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
)

type VoteAdminUseCase interface {
	// TODO update return type to int64
	Vote(ctx context.Context, businessId string, messageId string, userId string) (int64, error)

	UnVote(ctx context.Context, businessId string, messageId string, userId string) (int64, error)

	Count(ctx context.Context, businessId string, messageId string) (int64, error)

	IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error)

	VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error)

	CreateBusiness(ctx context.Context, business domain.Business) error
}
