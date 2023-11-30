package out

import "context"

type BusinessUseCase interface {
	CreateBusiness(ctx context.Context) error
	Business(ctx context.Context, businessId string) error
	Businesses(ctx context.Context) ([]string, error)
	UpdateBusiness(ctx context.Context, businessId string, businessName string) error
	DeleteBusiness(ctx context.Context, businessId string) error
}
