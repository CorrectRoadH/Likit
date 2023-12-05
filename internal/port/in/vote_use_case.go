package in

import "context"

type VoteUseCase interface {
	// TODO update return type to int64
	Vote(ctx context.Context, businessId string, messageId string, userId string) (int, error)

	UnVote(ctx context.Context, businessId string, messageId string, userId string) (int, error)

	Count(ctx context.Context, businessId string, messageId string) (int, error)

	IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error)

	VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error)
}
