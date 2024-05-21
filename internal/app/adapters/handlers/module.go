package handlers

import (
	"go.uber.org/fx"
	"golang-hexagon-example/internal/app/adapters/handlers/kafka"
)

var Module = fx.Options(
	kafka.Module,
)
