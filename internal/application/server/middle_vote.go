package server

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
)

// MiddleVoteServer is suit for middle vote server
type MiddleVoteServer struct {
}

func (m *MiddleVoteServer) Count(ctx context.Context, businessId string, messageId string) (int64, error) {
	panic("TODO: Implement")
}

func (m *MiddleVoteServer) IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error) {
	panic("TODO: Implement")
}

func (m *MiddleVoteServer) UnVote(ctx context.Context, businessId string, messageId string, userId string) (int64, error) {
	panic("TODO: Implement")
}

func (m *MiddleVoteServer) Vote(ctx context.Context, businessId string, messageId string, userId string) (int64, error) {
	// vote to redis
	// record vote event to postgres
	panic("TODO: Implement")
}

func (m *MiddleVoteServer) VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error) {
	panic("TODO: Implement")
}

func (m *MiddleVoteServer) Check(ctx context.Context, config domain.Config) error {
	// TODO check the whether have permission to create the table
	panic("TODO: Implement")
}

func (m *MiddleVoteServer) Features(ctx context.Context) {
	panic("TODO: Implement")
}

func (m *MiddleVoteServer) Init(ctx context.Context) {
	panic("TODO: Implement")
}

func (m *MiddleVoteServer) Requires(ctx context.Context) {
	panic("TODO: Implement")
}
