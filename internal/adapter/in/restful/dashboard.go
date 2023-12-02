package restful

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type DashboardServer struct {
}

func NewDashboardServer() *DashboardServer {
	s := &DashboardServer{}
	return s
}

//go:embed dist
var embeddedFiles embed.FS

func getFileSystem(path string) http.FileSystem {
	fs, err := fs.Sub(embeddedFiles, path)
	if err != nil {
		panic(err)
	}

	return http.FS(fs)
}

func (v *DashboardServer) register(e *echo.Echo) error {
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Request().URL.Path, "/api") || strings.HasPrefix(c.Request().URL.Path, "/admin") {
				return true
			}
			return false
		},
		HTML5:      true,
		Filesystem: getFileSystem("dist"),
	}))

	return nil
}
