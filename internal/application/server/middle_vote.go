package server

import (
	"context"
)

// MiddleVoteServer is suit for middle vote server
type MiddleVoteServer struct {
}

func (m *MiddleVoteServer) Count(ctx context.Context, businessId string, messageId string) (int, error) {
	panic("TODO: Implement")
}

func (m *MiddleVoteServer) IsVoted(ctx context.Context, businessId string, messageId string, userId string) (bool, error) {
	panic("TODO: Implement")
}

func (m *MiddleVoteServer) UnVote(ctx context.Context, businessId string, messageId string, userId string) error {
	panic("TODO: Implement")
}

func (m *MiddleVoteServer) Vote(ctx context.Context, businessId string, messageId string, userId string) error {
	panic("TODO: Implement")
}

func (m *MiddleVoteServer) VotedUsers(ctx context.Context, businessId string, messageId string) ([]string, error) {
	panic("TODO: Implement")
}
