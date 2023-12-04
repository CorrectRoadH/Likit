package in

import "github.com/CorrectRoadH/Likit/internal/application/domain"

type DatabaseUseCase interface {
	TestConnection(config domain.DatabaseConnectConfig) error
}
