// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v1/vote.proto

package apiv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	VoteService_Vote_FullMethodName       = "/likit.api.v1.VoteService/Vote"
	VoteService_UnVote_FullMethodName     = "/likit.api.v1.VoteService/UnVote"
	VoteService_Count_FullMethodName      = "/likit.api.v1.VoteService/Count"
	VoteService_IsVoted_FullMethodName    = "/likit.api.v1.VoteService/IsVoted"
	VoteService_VotedUsers_FullMethodName = "/likit.api.v1.VoteService/VotedUsers"
)

// VoteServiceClient is the client API for VoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VoteServiceClient interface {
	Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	UnVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error)
	IsVoted(ctx context.Context, in *IsVotedRequest, opts ...grpc.CallOption) (*IsVotedResponse, error)
	VotedUsers(ctx context.Context, in *VotedUsersRequest, opts ...grpc.CallOption) (*VotedUsersResponse, error)
}

type voteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVoteServiceClient(cc grpc.ClientConnInterface) VoteServiceClient {
	return &voteServiceClient{cc}
}

func (c *voteServiceClient) Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, VoteService_Vote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteServiceClient) UnVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, VoteService_UnVote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteServiceClient) Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, VoteService_Count_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteServiceClient) IsVoted(ctx context.Context, in *IsVotedRequest, opts ...grpc.CallOption) (*IsVotedResponse, error) {
	out := new(IsVotedResponse)
	err := c.cc.Invoke(ctx, VoteService_IsVoted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteServiceClient) VotedUsers(ctx context.Context, in *VotedUsersRequest, opts ...grpc.CallOption) (*VotedUsersResponse, error) {
	out := new(VotedUsersResponse)
	err := c.cc.Invoke(ctx, VoteService_VotedUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VoteServiceServer is the server API for VoteService service.
// All implementations must embed UnimplementedVoteServiceServer
// for forward compatibility
type VoteServiceServer interface {
	Vote(context.Context, *VoteRequest) (*VoteResponse, error)
	UnVote(context.Context, *VoteRequest) (*VoteResponse, error)
	Count(context.Context, *CountRequest) (*CountResponse, error)
	IsVoted(context.Context, *IsVotedRequest) (*IsVotedResponse, error)
	VotedUsers(context.Context, *VotedUsersRequest) (*VotedUsersResponse, error)
	mustEmbedUnimplementedVoteServiceServer()
}

// UnimplementedVoteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVoteServiceServer struct {
}

func (UnimplementedVoteServiceServer) Vote(context.Context, *VoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (UnimplementedVoteServiceServer) UnVote(context.Context, *VoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnVote not implemented")
}
func (UnimplementedVoteServiceServer) Count(context.Context, *CountRequest) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedVoteServiceServer) IsVoted(context.Context, *IsVotedRequest) (*IsVotedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsVoted not implemented")
}
func (UnimplementedVoteServiceServer) VotedUsers(context.Context, *VotedUsersRequest) (*VotedUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VotedUsers not implemented")
}
func (UnimplementedVoteServiceServer) mustEmbedUnimplementedVoteServiceServer() {}

// UnsafeVoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VoteServiceServer will
// result in compilation errors.
type UnsafeVoteServiceServer interface {
	mustEmbedUnimplementedVoteServiceServer()
}

func RegisterVoteServiceServer(s grpc.ServiceRegistrar, srv VoteServiceServer) {
	s.RegisterService(&VoteService_ServiceDesc, srv)
}

func _VoteService_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoteService_Vote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).Vote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoteService_UnVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).UnVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoteService_UnVote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).UnVote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoteService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoteService_Count_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).Count(ctx, req.(*CountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoteService_IsVoted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsVotedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).IsVoted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoteService_IsVoted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).IsVoted(ctx, req.(*IsVotedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoteService_VotedUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VotedUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).VotedUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoteService_VotedUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).VotedUsers(ctx, req.(*VotedUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VoteService_ServiceDesc is the grpc.ServiceDesc for VoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "likit.api.v1.VoteService",
	HandlerType: (*VoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Vote",
			Handler:    _VoteService_Vote_Handler,
		},
		{
			MethodName: "UnVote",
			Handler:    _VoteService_UnVote_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _VoteService_Count_Handler,
		},
		{
			MethodName: "IsVoted",
			Handler:    _VoteService_IsVoted_Handler,
		},
		{
			MethodName: "VotedUsers",
			Handler:    _VoteService_VotedUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/vote.proto",
}