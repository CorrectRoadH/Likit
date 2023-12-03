package server

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
	"github.com/CorrectRoadH/Likit/internal/port/in"
	"github.com/CorrectRoadH/Likit/internal/port/out"
)

type AdminServer struct {
	businessStore out.BusinessUseCase
}

func NewAdminServer(businessUseCase out.BusinessUseCase, redisConfig domain.RedisConfig) in.AdminUseCase {
	// init database
	business, err := businessUseCase.Businesses(context.Background())
	if err != nil {
		panic(err)
	}
	if len(business) == 0 {

		redisConfigJSON, err := domain.MarshalRedisConfig(redisConfig)
		if err != nil {
			panic(err)
		}
		businessUseCase.CreateBusiness(context.Background(), domain.Business{
			Title: "Comment Like",
			Id:    "COMMENT_LIKE",
			Type:  "SIMPLE_VOTE",
			Config: domain.Config{
				DataSourceConfig: []string{redisConfigJSON},
			},
		})
	}

	return &AdminServer{
		businessStore: businessUseCase,
	}
}

func (a *AdminServer) Businesses(ctx context.Context) ([]domain.Business, error) {
	return a.businessStore.Businesses(ctx)
}

func (a *AdminServer) CreateBusiness(ctx context.Context, business domain.Business) error {
	return a.businessStore.CreateBusiness(ctx, business)
}
