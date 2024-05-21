package core

import (
	"go.uber.org/fx"
	"golang-hexagon-example/internal/app/core/ports"
	"golang-hexagon-example/internal/app/core/services"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(services.NewUrlService, fx.As(new(ports.UrlService))),
		fx.Annotate(services.NewScraper, fx.As(new(ports.ScraperService))),
	),
)
