package out

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
)

type UserPort interface {
	CreateUser(ctx context.Context, user domain.User) error

	User(ctx context.Context, username string) (domain.User, error)

	Users(ctx context.Context) ([]domain.User, error)
}
