package out

import (
	"context"
	"github.com/CorrectRoadH/Likit/internal/application/domain"
)

type SaveUserUseCase interface {
	Login(ctx context.Context, username string, password string) (bool, error)

	CreateUser(ctx context.Context, username string, password string) error

	CreateUserByEnv(username string, password string) (domain.User, error)
}
