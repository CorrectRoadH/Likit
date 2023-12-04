package cmd

import (
	"github.com/CorrectRoadH/Likit/config"
	"github.com/CorrectRoadH/Likit/internal/adapter/in/restful"
	"github.com/CorrectRoadH/Likit/internal/adapter/in/restful/route"
	"github.com/CorrectRoadH/Likit/internal/adapter/out/admin"
	"github.com/CorrectRoadH/Likit/internal/adapter/out/database"
	"github.com/CorrectRoadH/Likit/internal/application/server"
	"go.uber.org/fx"
)

func Main() {
	fx.New(
		fx.Provide(
			restful.NewRESTfulServer,
			restful.NewVoteServer,
			restful.NewDashboardServer,

			route.NewAdminApiService,

			server.NewVoteServer,
			server.NewAdminServer,
			server.NewDatabaseServer,

			admin.NewBusinessAdapter,

			config.ProductEnvRedisConfig,
			config.ProductEnvConfigDatabaseConfig,

			database.NewRedisAdapter,
			database.NewPostgresAdapter,
		),
		fx.Invoke(func(s *restful.RESTfulServer) {
			s.Start()
		}),
	).Run()

}
