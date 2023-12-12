package in

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
)

type DatabaseUseCase interface {
	TestConnection(config domain.DatabaseConnectConfig) error

	DatabaseConfigureList() ([]domain.DatabaseConnectConfig, error)
	CreateDatabase(ctx context.Context, config domain.DatabaseConnectConfig) error
	UpdateDatabase(ctx context.Context, config domain.DatabaseConnectConfig) error
	DeleteDatabase(ctx context.Context, config domain.DatabaseConnectConfig) error
}
