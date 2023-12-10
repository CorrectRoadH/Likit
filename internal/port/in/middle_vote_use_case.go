package in

import "context"

type MiddleVoteUseCase interface {
	Rank(ctx context.Context, businessId string, messageId string)
}
