package main

import (
	"context"
	"poc-go-etl-consumer/configs"
	etlconsumerservice "poc-go-etl-consumer/pkg/etl_consumer"
	etlhandler "poc-go-etl-consumer/pkg/etl_consumer"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"golang.org/x/exp/slog"
)

func NewFiber(lc fx.Lifecycle, config *configs.Configs, etl *etlhandler.Handler) *fiber.App {

	port := ":3000"
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Mount("/etl", etl.App)

	slog.Info("Server is running on port " + port)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Listen(port)
			return nil
		},

		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
	return app
}

func main() {
	fx.New(

		fx.Provide(etlconsumerservice.NewHandler),
		fx.Provide(configs.NewConfigs),
		fx.Invoke(NewFiber),
	).Run()

}
