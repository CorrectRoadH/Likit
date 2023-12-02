package route

import (
	"net/url"
	"strings"

	"github.com/CorrectRoadH/Likit/codegen"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

var (
	_swagger *openapi3.T

	V2APIPath string
)

type api struct{}

func init() {
	swagger, err := codegen.GetSwagger()
	if err != nil {
		panic(err)
	}

	_swagger = swagger

	u, err := url.Parse(_swagger.Servers[0].URL)
	if err != nil {
		panic(err)
	}

	V2APIPath = strings.TrimRight(u.Path, "/")
}

func NewAPIService() codegen.ServerInterface {
	return &api{}
}

func (a *api) Login(ctx echo.Context) error {
	panic("TODO: Implement")
}
