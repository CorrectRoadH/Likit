package out

import "github.com/CorrectRoadH/Likit/internal/application/domain"

type RedisPort interface {
	TestConnect(config domain.DatabaseConnectConfig) error
}
