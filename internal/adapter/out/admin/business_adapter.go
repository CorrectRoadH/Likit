package admin

import (
	"database/sql"

	"context"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/out"
)

type BusinessAdapter struct {
	db *sql.DB
}

func (b *BusinessAdapter) Business(ctx context.Context, businessId string) error {
	panic("TODO: Implement")
}

func (b *BusinessAdapter) Businesses(ctx context.Context) ([]string, error) {
	panic("TODO: Implement")
}

func (b *BusinessAdapter) CreateBusiness(ctx context.Context) error {
	panic("TODO: Implement")
}

func (b *BusinessAdapter) DeleteBusiness(ctx context.Context, businessId string) error {
	panic("TODO: Implement")
}

func (b *BusinessAdapter) UpdateBusiness(ctx context.Context, businessId string, businessName string) error {
	panic("TODO: Implement")
}

func NewBusinessAdapter(config domain.ConfigDatabaseConfig) out.BusinessUseCase {
	db, err := sql.Open("postgres", string(config))
	if err != nil {
		panic(err)
	}

	return &BusinessAdapter{
		db: db,
	}
}
