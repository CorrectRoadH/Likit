package out

import "github.com/CorrectRoadH/Likit/internal/application/domain"

type DatabasePort interface {
	CreateDatabaseConnectConfig(config domain.DatabaseConnectConfig) error
	DatabaseConnectConfig(configId string) (domain.DatabaseConnectConfig, error)
	DeleteDatabaseConnectConfig(configId string) error
	ListDatabaseConnectConfig() ([]domain.DatabaseConnectConfig, error)
}
