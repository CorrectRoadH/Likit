package server

import (
	"context"

	"github.com/CorrectRoadH/Likit/internal/application/domain"
)

type AdminServer struct {
}

func (a *AdminServer) Business(ctx context.Context, business domain.Business) error {
	panic("TODO: Implement")
}

func (a *AdminServer) CreateBusiness(ctx context.Context) error {
	panic("TODO: Implement")
}
