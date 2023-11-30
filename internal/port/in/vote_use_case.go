package in

import "context"

type VoteUseCase interface {
	Vote(ctx context.Context, businessId string, messageId string, userId string) error

	UnVote(ctx context.Context, businessId string, messageId string, userId string) error

	Count(ctx context.Context, businessId string, messageId string) (int, error)

	IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error)

	UserListOfVoted(ctx context.Context, businessId string, messageId string) ([]string, error)
}
