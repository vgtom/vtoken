package bootstrap

import (
	"context"

	"github.com/vgtom/vtoken/api/controller"
	"github.com/vgtom/vtoken/api/repository"
	"github.com/vgtom/vtoken/api/routes"
	"github.com/vgtom/vtoken/api/service"
	"github.com/vgtom/vtoken/cron"
	"github.com/vgtom/vtoken/lib"
	"github.com/vgtom/vtoken/middlewares"

	"go.uber.org/fx"
)

var Module = fx.Options(
	lib.Module,
	middlewares.Module,
	controller.Module,
	routes.Module,
	service.Module,
	repository.Module,
	cron.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routes.Routes,
	env lib.Env,
	database lib.Database,
) {
	conn, _ := database.DB.DB()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			conn.SetMaxOpenConns(10)
			go func() {
				routes.Setup()
				handler.Gin.Run(":" + env.ServerPort)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			conn.Close()
			return nil
		},
	})
}
