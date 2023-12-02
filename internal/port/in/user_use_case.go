package in

import "context"

type UserUseCase interface {
	Login(ctx context.Context, username string, password string) (bool, error)

	CreateUser(ctx context.Context, username string, password string) error

	CreateUserByEnv(username string, password string) error
}
