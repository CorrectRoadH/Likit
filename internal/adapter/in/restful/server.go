package restful

import (
	"github.com/CorrectRoadH/Likit/codegen"
	"github.com/labstack/echo/v4"
)

type RESTfulServer struct {
	e *echo.Echo

	voteServer *VoteServer
}

func NewRESTfulServer(
	voteServer *VoteServer,
	dashboardServer *DashboardServer,

	apiService codegen.ServerInterface,
) *RESTfulServer {
	e := echo.New()

	// TODO add a middleware to return 500 err for vote no exist business id

	s := &RESTfulServer{
		e:          e,
		voteServer: voteServer,
	}

	dashboardServer.register(e)

	voteServer.register(e.Group("/api/v1"))

	codegen.RegisterHandlersWithBaseURL(e, apiService, "/admin/v1")

	return s
}

func (v *RESTfulServer) Start() error {

	v.e.Start(":7789")
	return nil
}
