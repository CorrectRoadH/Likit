package grpc

import (
	"context"

	apiv1 "github.com/CorrectRoadH/Likit/grpc/gen/api/v1"
	"github.com/CorrectRoadH/Likit/internal/port/in"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VoteGRPCServer struct {
	apiv1.UnimplementedVoteServiceServer
	voteUseCase in.VoteAdminUseCase
}

func NewVoteGRPCServer(voteUseCase in.VoteAdminUseCase) *VoteGRPCServer {
	s := &VoteGRPCServer{
		voteUseCase: voteUseCase,
	}
	return s
}

func (s *VoteGRPCServer) Vote(ctx context.Context, in *apiv1.VoteRequest) (*apiv1.VoteResponse, error) {
	val, err := s.voteUseCase.Vote(ctx, in.BusinessId, in.MessageId, in.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &apiv1.VoteResponse{
		Count: val,
	}, nil
}

func (s *VoteGRPCServer) UnVote(ctx context.Context, in *apiv1.VoteRequest) (*apiv1.VoteResponse, error) {
	val, err := s.voteUseCase.UnVote(ctx, in.BusinessId, in.MessageId, in.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &apiv1.VoteResponse{
		Count: val,
	}, nil
}

func (s *VoteGRPCServer) Count(ctx context.Context, in *apiv1.CountRequest) (*apiv1.CountResponse, error) {
	val, err := s.voteUseCase.Count(ctx, in.BusinessId, in.MessageId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &apiv1.CountResponse{
		Count: val,
	}, nil
}

func (s *VoteGRPCServer) IsVoted(ctx context.Context, in *apiv1.IsVotedRequest) (*apiv1.IsVotedResponse, error) {
	val, err := s.voteUseCase.IsVoted(ctx, in.BusinessId, in.MessageId, in.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &apiv1.IsVotedResponse{
		IsVoted: val,
	}, nil
}

func (s *VoteGRPCServer) VotedUsers(ctx context.Context, in *apiv1.VotedUsersRequest) (*apiv1.VotedUsersResponse, error) {
	val, err := s.voteUseCase.VotedUsers(ctx, in.BusinessId, in.MessageId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &apiv1.VotedUsersResponse{
		UserIds: val,
	}, nil
}
