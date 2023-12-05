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
	Id           string       `json:"id"`
	Title        string       `json:"title"`
	DatabaseType DatabaseType `json:"database_type"`
	Host         string       `json:"host"`
	Port         int          `json:"port"`
	Username     string       `json:"username"`
	Password     string       `json:"password"`
	Database     string       `json:"database"`
}

// func (c DatabaseConnectConfig) Value() (driver.Value, error) {
// 	str, err := json.Marshal(c)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return string(str), nil
// }

// func (c *DatabaseConnectConfig) Scan(value interface{}) error {
// 	str, ok := value.(string)
// 	if !ok {
// 		return errors.New("数据类型不匹配")
// 	}

// 	return json.Unmarshal([]byte(str), c)
// }

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
