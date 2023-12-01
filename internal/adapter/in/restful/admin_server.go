package restful

import (
	"github.com/labstack/echo/v4"
)

type AdminServer struct {
}

func NewAdminServer() *AdminServer {
	s := &AdminServer{}
	return s
}

func (v *AdminServer) register(g *echo.Group) error {

	return nil
}
