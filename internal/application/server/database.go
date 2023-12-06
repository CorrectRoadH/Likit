package server

import (
	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/CorrectRoadH/Likit/internal/port/out"
)

type DatabaseServer struct {
	databaseServer out.DatabasePort

	postgres out.PostgresPort
	redis    out.RedisPort
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
	postgres out.PostgresPort,
	redis out.RedisPort,

	databaseServer out.DatabasePort,

	defaultPostgresConfig domain.PostgresConfig,
	defaultRedisConfig domain.RedisConfig,
) in.DatabaseUseCase {
	// init database
	configs, err := databaseServer.ListDatabaseConnectConfig()
	if err != nil {
		panic(err)
	}
	if len(configs) == 0 {
		err := databaseServer.CreateDatabaseConnectConfig(domain.DatabaseConnectConfig(defaultPostgresConfig))
		if err != nil {
			panic(err)
		}
		err = databaseServer.CreateDatabaseConnectConfig(domain.DatabaseConnectConfig(defaultRedisConfig))
		if err != nil {
			panic(err)
		}
	}

	return &DatabaseServer{
		postgres:       postgres,
		redis:          redis,
		databaseServer: databaseServer,
	}
}
