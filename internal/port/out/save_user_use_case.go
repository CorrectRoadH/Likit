package out

import "context"

type SaveUserUseCase interface {
	Login(ctx context.Context, username string, password string) (bool, error)

	CreateUser(ctx context.Context, username string, password string) error

	CreateUserByEnv(username string, password string) error
}
