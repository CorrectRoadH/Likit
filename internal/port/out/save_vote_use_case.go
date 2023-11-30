package out

import "context"

type SaveVoteUseCase interface {
	Vote(ctx context.Context, businessId string, messageId string, userId string) error
	UnVote(ctx context.Context, businessId string, messageId string, userId string) error

	Count(ctx context.Context, businessId string, messageId string) (int, error)

	VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error)

	IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error)
}
