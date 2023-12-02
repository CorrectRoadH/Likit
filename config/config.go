package config

import (
	"os"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
)

func ProductEnvRedisConfig() domain.RedisConfig {
	addr := os.Getenv("REDIS_URI") // "localhost:6379"
	passwd := os.Getenv("PASSWORD")

	return domain.RedisConfig{
		Addr:   addr,
		Passwd: passwd,
	}
}

func TestEnvRedisConfig() domain.RedisConfig {
	addr := os.Getenv("REDIS_URI") // "localhost:6379"
	passwd := os.Getenv("PASSWORD")

	return domain.RedisConfig{
		Addr:   addr,
		Passwd: passwd,
	}
}

// the config is for business data. not vote data
func ProductEnvConfigDatabaseConfig() domain.ConfigDatabaseConfig {
	// addr := os.Getenv("REDIS_URI") // "localhost:6379"
	// passwd := os.Getenv("PASSWORD")

	return domain.ConfigDatabaseConfig("postgres://postgres:postgres@localhost:5432/likit?sslmode=disable")
}
