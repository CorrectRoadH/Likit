package in

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
)

type BasicVoteUseCase interface {
	Check(ctx context.Context, config domain.Config) error
	ID(ctx context.Context) string
	Features(ctx context.Context)
	Requires(ctx context.Context)
	Init(ctx context.Context)
}
