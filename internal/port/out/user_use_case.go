package out

import (
	"context"
	"github.com/CorrectRoadH/Likit/internal/application/domain"
)

type UserPortUseCase interface {
	CreateUser(ctx context.Context, username string, password string) error

	CreateUserByEnv() (domain.User, error)
}
