package server

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
)

// MiddleVoteEngine is suit for middle vote server
type MiddleVoteEngine struct {
}

func (m *MiddleVoteEngine) Count(ctx context.Context, businessId string, messageId string) (int64, error) {
	panic("TODO: Implement")
}

func (m *MiddleVoteEngine) IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error) {
	panic("TODO: Implement")
}

func (m *MiddleVoteEngine) UnVote(ctx context.Context, businessId string, messageId string, userId string) (int64, error) {
	panic("TODO: Implement")
}

func (m *MiddleVoteEngine) Vote(ctx context.Context, businessId string, messageId string, userId string) (int64, error) {
	// vote to redis
	// record vote event to postgres
	panic("TODO: Implement")
}

func (m *MiddleVoteEngine) VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error) {
	panic("TODO: Implement")
}

func (m *MiddleVoteEngine) Check(ctx context.Context, config domain.Config) error {
	// TODO check the whether have permission to create the table
	panic("TODO: Implement")
}

func (m *MiddleVoteEngine) Features(ctx context.Context) {
	panic("TODO: Implement")
}

func (m *MiddleVoteEngine) Init(ctx context.Context) {
	panic("TODO: Implement")
}

func (m *MiddleVoteEngine) Requires(ctx context.Context) {
	panic("TODO: Implement")
}
