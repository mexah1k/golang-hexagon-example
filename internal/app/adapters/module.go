package adapters

import (
	"go.uber.org/fx"
	"golang-hexagon-example/internal/app/adapters/handlers"
	"golang-hexagon-example/internal/app/adapters/repositories"
)

var Module = fx.Options(
	handlers.Module,
	repositories.Module,
)
