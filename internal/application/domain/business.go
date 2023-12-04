package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Business struct {
	Title    string `json:"title"`
	Id       string `json:"business_id"`
	Type     string `json:"type"`
	CreateAt string
	UpdateAt string

	Config Config `json:"config"`
}

type Config struct {
	DataSourceConfig []DatabaseConnectConfig `json:"dataSourceConfig"`
}

// for gorm
func (c Config) Value() (driver.Value, error) {
	str, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return string(str), nil
}

func (c *Config) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("数据类型不匹配")
	}

	return json.Unmarshal([]byte(str), c)
}
