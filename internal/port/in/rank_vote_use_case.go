package in

import "context"

type RankVoteUseCase interface {
	Rank(ctx context.Context, businessId string, limit int) []string
	RankFromMessages(ctx context.Context, businessId string, messageIds []string) []string
}
