package cmd

import (
	"github.com/CorrectRoadH/Likit/config"
	"github.com/CorrectRoadH/Likit/internal/adapter/in/restful"
	v1 "github.com/CorrectRoadH/Likit/internal/adapter/out/v1"
	"github.com/CorrectRoadH/Likit/internal/application/server"
	"go.uber.org/fx"
)

func Main() {
	fx.New(
		fx.Provide(
			restful.NewRESTfulServer,
			restful.NewVoteServer,
			restful.NewAdminServer,
			restful.NewDashboardServer,

			server.NewVoteServer,
			v1.NewRedisAdapter,

			config.ProductEnvRedisConfig,
		),
		fx.Invoke(func(s *restful.RESTfulServer) {
			s.Start()
		}),
	).Run()

}
