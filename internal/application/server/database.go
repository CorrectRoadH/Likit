package server

import (
	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/CorrectRoadH/Likit/internal/port/out"
)

type DatabaseServer struct {
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

func NewDatabaseServer(
	postgres out.PostgresUseCase,
	redis out.RedisUseCase,
) in.DatabaseUseCase {

	return &DatabaseServer{
		postgres: postgres,
		redis:    redis,
	}
}
