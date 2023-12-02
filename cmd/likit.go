package cmd

import (
	"github.com/CorrectRoadH/Likit/config"
	"github.com/CorrectRoadH/Likit/internal/adapter/in/restful"
	"github.com/CorrectRoadH/Likit/internal/adapter/in/restful/route"
	"github.com/CorrectRoadH/Likit/internal/adapter/out/admin"
	"github.com/CorrectRoadH/Likit/internal/application/server"
	"go.uber.org/fx"
)

func Main() {
	fx.New(
		fx.Provide(
			restful.NewRESTfulServer,
			restful.NewVoteServer,
			restful.NewDashboardServer,
			restful.NewUserServer,

			route.NewAdminApiService,

			server.NewVoteServer,
			server.NewAdminServer,

			admin.NewBusinessAdapter,
			server.NewUserServer,

			config.ProductEnvRedisConfig,
			config.ProductEnvConfigDatabaseConfig,
		),
		fx.Invoke(func(s *restful.RESTfulServer) {
			s.Start()
		}),
	).Run()

}
