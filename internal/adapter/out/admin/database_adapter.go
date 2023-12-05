package admin

import (
	"fmt"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/out"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseAdapter struct {
	db *gorm.DB
}

func NewDatabaseAdapter(config domain.PostgresConfig) out.DatabaseUseCase {
	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.Username, config.Password, config.Host, config.Port, config.Database),
		), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	db.AutoMigrate(&domain.DatabaseConnectConfig{})
	return &DatabaseAdapter{
		db: db,
	}
}

func (d *DatabaseAdapter) CreateDatabaseConnectConfig(config domain.DatabaseConnectConfig) error {
	return d.db.Create(&config).Error
}

func (d *DatabaseAdapter) DatabaseConnectConfig(configId string) (domain.DatabaseConnectConfig, error) {
	var config domain.DatabaseConnectConfig

	result := d.db.First(&config, "id = ?", configId)
	return config, result.Error
}

func (d *DatabaseAdapter) DeleteDatabaseConnectConfig(configId string) error {
	result := d.db.Delete(&domain.DatabaseConnectConfig{}, "id = ?", configId)
	return result.Error
}

func (d *DatabaseAdapter) ListDatabaseConnectConfig() ([]domain.DatabaseConnectConfig, error) {
	var configs []domain.DatabaseConnectConfig

	result := d.db.Find(&configs)
	return configs, result.Error
}
