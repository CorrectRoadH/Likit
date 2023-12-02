package restful

import (
	"fmt"
	"os"

	"github.com/CorrectRoadH/Likit/codegen"
	"github.com/CorrectRoadH/Likit/internal/adapter/in/restful/route"
	"github.com/labstack/echo/v4"
)

type RESTfulServer struct {
	e *echo.Echo

	voteServer  *VoteServer
	adminServer *AdminServer
}

func NewRESTfulServer(
	voteServer *VoteServer,
	adminServer *AdminServer,
	dashboardServer *DashboardServer,

	apiService codegen.ServerInterface,
) *RESTfulServer {
	e := echo.New()

	// TODO add a middleware to return 500 err for vote no exist business id

	s := &RESTfulServer{
		e:           e,
		voteServer:  voteServer,
		adminServer: adminServer,
	}

	dashboardServer.register(e)

	voteServer.register(e.Group("/api/v1"))
	// adminServer.register(e.Group("/admin/v1"))

	codegen.RegisterHandlersWithBaseURL(e, apiService, route.V2APIPath)

	return s
}

func (v *RESTfulServer) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	v.e.Start(fmt.Sprintf(":%s", port))
	return nil
}
