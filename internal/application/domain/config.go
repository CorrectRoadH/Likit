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
	DatabaseName string       `json:"database_name"`
}

type RedisConfig DatabaseConnectConfig
type PostgresConfig DatabaseConnectConfig

// type RedisConfig struct {
// 	Addr   string `json:"addr"`
// 	Passwd string `json:"passwd"`
// }

// func UnmarshalRedisConfig(config string) (RedisConfig, error) {
// 	if strings.HasPrefix(config, "redisconfig:") {
// 		redisConfigJson := config[12:]
// 		var redisConfig RedisConfig
// 		err := json.Unmarshal([]byte(redisConfigJson), &redisConfig)
// 		if err != nil {
// 			return RedisConfig{}, err
// 		}
// 		return redisConfig, nil
// 	}
// 	return RedisConfig{}, fmt.Errorf("no redis config")
// }

// func MarshalRedisConfig(config RedisConfig) (string, error) {
// 	configJson, err := json.Marshal(config)

// 	return "redisconfig:" + string(configJson), err
// }

func GetRedisConfig(config Config) (RedisConfig, error) {
	for _, config := range config.DataSourceConfig {
		if config.DatabaseType != REDIS {
			continue
		}
		return RedisConfig(config), nil
	}
	return RedisConfig{}, fmt.Errorf("no redis config in config")
}
