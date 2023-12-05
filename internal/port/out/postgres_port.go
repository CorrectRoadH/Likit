package out

import "github.com/CorrectRoadH/Likit/internal/application/domain"

type PostgresPort interface {
	TestConnect(config domain.DatabaseConnectConfig) error
}
