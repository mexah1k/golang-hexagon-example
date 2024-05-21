package repositories

import (
	"go.uber.org/fx"
	"golang-hexagon-example/internal/app/adapters/repositories/mongodb"
)

var Module = fx.Options(
	mongodb.Module,
)
