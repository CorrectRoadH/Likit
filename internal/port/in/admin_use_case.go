package in

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
)

type AdminUseCase interface {
	CreateBusiness(ctx context.Context, business domain.Business) error
	Businesses(ctx context.Context) ([]domain.Business, error)
}
