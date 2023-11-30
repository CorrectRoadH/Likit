package admin

import (
	"context"
	"fmt"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/out"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type BusinessAdapter struct {
	db *gorm.DB
}

func (b *BusinessAdapter) Business(ctx context.Context, businessId string) (domain.Business, error) {
	var business domain.Business

	result := b.db.First(&business, "id = ?", businessId)
	return business, result.Error
}

func (b *BusinessAdapter) Businesses(ctx context.Context) ([]domain.Business, error) {
	var businesses []domain.Business

	result := b.db.Find(&businesses)
	return businesses, result.Error
}

func (b *BusinessAdapter) DeleteBusiness(ctx context.Context, businessId string) error {

	result := b.db.Delete(&domain.Business{}, "id = ?", businessId)
	return result.Error
}

func (b *BusinessAdapter) UpdateBusiness(ctx context.Context, businessId string, business domain.Business) error {
	panic("TODO: Implement")
}

func (b *BusinessAdapter) CreateBusiness(business domain.Business, ctx context.Context) error {
	result := b.db.Create(&business)
	return result.Error
}

func NewBusinessAdapter(config domain.ConfigDatabaseConfig) out.BusinessUseCase {
	db, err := gorm.Open(postgres.Open(string(config)), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	db.AutoMigrate(&domain.Business{})
	return &BusinessAdapter{
		db: db,
	}
}
