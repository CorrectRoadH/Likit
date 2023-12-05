package server

import (
	"context"
	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/CorrectRoadH/Likit/internal/port/out"
)

type UserServer struct {
	userStore out.UserPortUseCase
}

func NewUserServer(userStore out.UserPortUseCase) in.UserUseCase {
	return &UserServer{
		userStore: userStore,
	}
}

func (u UserServer) Login(ctx context.Context, username string, password string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServer) CreateUser(ctx context.Context, username string, password string) error {
	//TODO implement me
	panic("implement me")
}
