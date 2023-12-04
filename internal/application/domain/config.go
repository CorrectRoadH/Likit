package domain

import (
	"fmt"
)

type DatabaseType string

const (
	REDIS    DatabaseType = "redis"
	POSTGRES DatabaseType = "postgres"
)

type DatabaseConnectConfig struct {
	DatabaseType DatabaseType `json:"database_type"`
	Host         string       `json:"host"`
	Port         int          `json:"port"`
	Username     string       `json:"username"`
	Password     string       `json:"password"`
	Database     string       `json:"database"`
}

type RedisConfig DatabaseConnectConfig
type PostgresConfig DatabaseConnectConfig

func GetRedisConfig(config Config) (RedisConfig, error) {
	for _, config := range config.DataSourceConfig {
		if config.DatabaseType != REDIS {
			continue
		}
		return RedisConfig(config), nil
	}
	return RedisConfig{}, fmt.Errorf("no redis config in config")
}
