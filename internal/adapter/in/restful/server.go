package restful

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

type RESTfulServer struct {
	e *echo.Echo

	voteServer  *VoteServer
	adminServer *AdminServer
}

func NewRESTfulServer(voteServer *VoteServer, adminServer *AdminServer) *RESTfulServer {
	e := echo.New()

	s := &RESTfulServer{
		e:           e,
		voteServer:  voteServer,
		adminServer: adminServer,
	}

	voteServer.register(e.Group("/api/v1"))
	adminServer.register(e.Group("/admin/v1"))

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
