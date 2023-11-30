package restful

import "github.com/CorrectRoadH/Likit/internal/port/in"

type VoteServer struct {
	voteUseCase *in.VoteUseCase
}

func NewVoteServer(voteUseCase *in.VoteUseCase) *VoteServer {
	return &VoteServer{
		voteUseCase: voteUseCase,
	}
}

func (v *VoteServer) Start() error {
	return nil
}
