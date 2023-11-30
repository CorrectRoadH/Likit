package out

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
)

// TODO Change the name to a better one
type BusinessUseCase interface {
	CreateBusiness(business domain.Business, ctx context.Context) error
	Business(ctx context.Context, businessId string) (domain.Business, error)
	Businesses(ctx context.Context) ([]domain.Business, error)
	UpdateBusiness(ctx context.Context, businessId string, business domain.Business) error
	DeleteBusiness(ctx context.Context, businessId string) error
}
