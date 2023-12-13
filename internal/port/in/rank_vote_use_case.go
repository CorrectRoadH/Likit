package in

import "context"

type RankVoteUseCase interface {
	Rank(ctx context.Context, businessId string, messageId string)
}
