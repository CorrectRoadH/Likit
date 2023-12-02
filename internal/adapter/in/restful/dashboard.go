package restful

import (
	"embed"
	"io/fs"
	"net/http"

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
		HTML5:      false,
		Filesystem: getFileSystem("dist"),
	}))

	return nil
}
