package cmd

import (
	"github.com/CorrectRoadH/Likit/config"
	"github.com/CorrectRoadH/Likit/internal/adapter/in/grpc"
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

			grpc.NewGrpcServer,
			grpc.NewVoteGRPCServer,

			server.NewVoteServer,
			server.NewAdminServer,
			server.NewDatabaseServer,
			server.NewUserServer,

			admin.NewBusinessAdapter,
			admin.NewDatabaseAdapter,
			admin.NewUserAdapter,

			config.ProductEnvRedisConfig,
			config.ProductEnvConfigDatabaseConfig,

			database.NewRedisAdapter,
			database.NewPostgresAdapter,
		),
		fx.Invoke(func(s *restful.RESTfulServer, g *grpc.GrpcServer) {
			s.Start()
		}),
	).Run()

}
