package server

import (
	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/CorrectRoadH/Likit/internal/port/out"
)

type DatabaseServer struct {
	databaseServer out.DatabaseUseCase

	postgres out.PostgresUseCase
	redis    out.RedisUseCase
}

func (d *DatabaseServer) TestConnection(config domain.DatabaseConnectConfig) error {
	if config.DatabaseType == domain.REDIS {
		return d.redis.TestConnect(config)
	}

	if config.DatabaseType == domain.POSTGRES {
		return d.postgres.TestConnect(config)
	}
	return domain.ErrDatabaseTypeNotSupport
}

func (d *DatabaseServer) DatabaseConfigureList() ([]domain.DatabaseConnectConfig, error) {
	return d.databaseServer.ListDatabaseConnectConfig()
}

func NewDatabaseServer(
	postgres out.PostgresUseCase,
	redis out.RedisUseCase,

	databaseServer out.DatabaseUseCase,

	defaultPostgresConfig domain.PostgresConfig,
	defaultRedisConfig domain.RedisConfig,
) in.DatabaseUseCase {
	// init database
	configs, err := databaseServer.ListDatabaseConnectConfig()
	if err != nil {
		panic(err)
	}
	if len(configs) == 0 {
		databaseServer.CreateDatabaseConnectConfig(domain.DatabaseConnectConfig(defaultPostgresConfig))
		databaseServer.CreateDatabaseConnectConfig(domain.DatabaseConnectConfig(defaultRedisConfig))
	}

	return &DatabaseServer{
		postgres:       postgres,
		redis:          redis,
		databaseServer: databaseServer,
	}
}
