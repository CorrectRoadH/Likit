package out

import "github.com/CorrectRoadH/Likit/internal/application/domain"

type PostgresUseCase interface {
	TestConnect(config domain.DatabaseConnectConfig) error
}
