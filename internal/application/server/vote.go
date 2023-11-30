package server

import "github.com/CorrectRoadH/Likit/internal/port/out"

type VoteServer struct {
	VoteStore out.SaveVoteUseCase
}

func (v *VoteServer) Vote(businessId string, messageId string, userId string) error {
	return nil
}

func (v *VoteServer) Count(businessId, messageId, userId string) (int, error) {
	return 0, nil
}
