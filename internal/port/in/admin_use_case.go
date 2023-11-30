package in

import "context"

type AdminUseCase interface {
	CreateBusiness(ctx context.Context) error
	Business(ctx context.Context) error
}
