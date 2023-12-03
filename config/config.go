package config

import (
	"fmt"
	"os"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
)

func ProductEnvRedisConfig() domain.RedisConfig {
	addr := os.Getenv("REDIS_URI") // "localhost:6379"
	passwd := os.Getenv("PASSWORD")

	fmt.Println("redis uri: ", addr)
	if addr == "" {
		addr = "localhost:6379"
	}

	return domain.RedisConfig{
		Addr:   addr,
		Passwd: passwd,
	}
}

func TestEnvRedisConfig() domain.RedisConfig {
	addr := os.Getenv("REDIS_URI") // "localhost:6379"
	passwd := os.Getenv("PASSWORD")

	if addr == "" {
		addr = "localhost:6379"
	}

	return domain.RedisConfig{
		Addr:   addr,
		Passwd: passwd,
	}
}

// the config is for business data. not vote data
func ProductEnvConfigDatabaseConfig() domain.ConfigDatabaseConfig {
	connectString := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connectString == "" {
		connectString = "postgres://postgres:postgres@localhost:5432/likit?sslmode=disable"
	}

	fmt.Println("postgres uri: ", connectString)

	return domain.ConfigDatabaseConfig(connectString)
}
