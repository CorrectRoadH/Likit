package out

import "github.com/CorrectRoadH/Likit/internal/application/domain"

type RedisUseCase interface {
	TestConnect(config domain.DatabaseConnectConfig) error
}
