package app

import (
	"go.uber.org/fx"
	"golang-hexagon-example/internal/app/adapters"
	"golang-hexagon-example/internal/app/core"
	"golang-hexagon-example/internal/app/infrastructure/config"
	"golang-hexagon-example/internal/app/infrastructure/server"
)

func NewApp(cfg *config.AppConfig) *fx.App {
	return fx.New(
		fx.Provide(func() *config.AppConfig {
			return cfg
		}),
		adapters.Module,
		core.Module,
		server.Module,
		fx.Invoke(
			server.StartServer,
		),
	)
}
