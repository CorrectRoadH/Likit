package out

import "context"

type SaveVoteUseCase interface {
	Vote(ctx context.Context, businessId string, messageId string, userId string) error

	Count(ctx context.Context, businessId string, messageId string) (int, error)
}
