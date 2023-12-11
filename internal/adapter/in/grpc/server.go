package grpc

import (
	"log"
	"net"
	"os"

	apiv1 "github.com/CorrectRoadH/Likit/grpc/gen/api/v1"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	listener net.Listener
	Server   *grpc.Server
}

func NewGrpcServer(voteGRPCServer *VoteGRPCServer) *GrpcServer {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "4778"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	} else {
		log.Println("GRPC Server is listening on port", port)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	apiv1.RegisterVoteServiceServer(s, voteGRPCServer)

	return &GrpcServer{
		Server:   s,
		listener: lis,
	}
}
