package config

import (
	"os"
	"strconv"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
)

func ProductEnvRedisConfig() domain.DatabaseConnectConfig {
	host := os.Getenv("REDIS_HOST") // "localhost:6379"
	port := os.Getenv("REDIS_PORT")
	passwd := os.Getenv("PASSWORD")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "6379"
	}

	// port to int
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	return domain.DatabaseConnectConfig{
		DatabaseType: domain.REDIS,
		Host:         host,
		Port:         portInt,
		Password:     passwd,
	}
}

func TestEnvRedisConfig() domain.DatabaseConnectConfig {
	host := os.Getenv("REDIS_HOST") // "localhost:6379"
	port := os.Getenv("REDIS_PORT")
	passwd := os.Getenv("PASSWORD")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "6379"
	}

	// port to int
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	return domain.DatabaseConnectConfig{
		DatabaseType: domain.REDIS,
		Host:         host,
		Port:         portInt,
		Password:     passwd,
	}
}

// the config is for business data. not vote data
func ProductEnvConfigDatabaseConfig() domain.DatabaseConnectConfig {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DATABASE_NAME")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if username == "" {
		username = "postgres"
	}
	if password == "" {
		password = "postgres"
	}
	if databaseName == "" {
		databaseName = "likit"
	}

	// port to int
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	return domain.DatabaseConnectConfig{
		DatabaseType: domain.POSTGRES,
		Host:         host,
		Port:         portInt,
		Username:     username,
		Password:     password,
		DatabaseName: databaseName,
	}
}
