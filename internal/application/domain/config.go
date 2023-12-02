package domain

import (
	"encoding/json"
	"fmt"
	"strings"
)

type RedisConfig struct {
	Addr   string `json:"addr"`
	Passwd string `json:"passwd"`
}

func UnmarshalRedisConfig(config string) (RedisConfig, error) {
	if strings.HasPrefix(config, "redisconfig:") {
		redisConfigJson := config[12:]
		var redisConfig RedisConfig
		err := json.Unmarshal([]byte(redisConfigJson), &redisConfig)
		if err != nil {
			return RedisConfig{}, err
		}
		return redisConfig, nil
	}
	return RedisConfig{}, fmt.Errorf("no redis config")
}

func MarshalRedisConfig(config RedisConfig) (string, error) {
	configJson, err := json.Marshal(config)

	return "redisconfig:" + string(configJson), err
}

func GetRedisConfig(config Config) (RedisConfig, error) {
	for _, config := range config.DataSourceConfig {
		config, err := UnmarshalRedisConfig(config)
		if err != nil {
			continue
		}
		return config, nil
	}
	return RedisConfig{}, fmt.Errorf("no redis config in config")
}

type ConfigDatabaseConfig string
