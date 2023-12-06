package server

import (
	"context"
	"os"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/CorrectRoadH/Likit/internal/port/out"
	"github.com/CorrectRoadH/Likit/utils"
)

type UserServer struct {
	userStore out.UserPort
}

func NewUserServer(userStore out.UserPort) in.UserUseCase {
	// init database
	users, err := userStore.Users(context.Background())
	if err != nil {
		panic(err)
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	if username == "" {
		username = "admin"
	}
	if password == "" {
		password = "admin"
	}

	if len(users) == 0 {
		userStore.CreateUser(context.Background(), domain.User{
			Id:       utils.Uuid(),
			Username: username,
			Password: password,
		})
	}

	return &UserServer{
		userStore: userStore,
	}
}

func (u UserServer) Login(ctx context.Context, username string, password string) (bool, error) {
	user, err := u.userStore.User(ctx, username)
	if err != nil {
		return false, err
	}
	return user.Password == password, nil
}

func (u UserServer) CreateUser(ctx context.Context, username string, password string) error {
	return u.userStore.CreateUser(ctx, domain.User{
		Id:       utils.Uuid(),
		Username: username,
		Password: password,
	})
}
