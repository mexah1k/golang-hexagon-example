package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
	"golang-hexagon-example/internal/app/core/ports"
	"golang-hexagon-example/internal/app/infrastructure/config"
)

var Module = fx.Options(
	fx.Provide(
		func(config *config.AppConfig) MongoDBConfig {
			conn := config.Mongo.ConnectionString
			if conn == "" && config.Environment == "Development" { // for local development
				conn = "mongodb://localhost:27017"
			}

			return MongoDBConfig{URI: conn}
		},
		func(cfg MongoDBConfig) (*mongo.Client, error) {
			return ConnectToMongo(cfg)
		},
	),
	fx.Provide(
		fx.Annotate(NewUrlRepository, fx.As(new(ports.UrlRepository))),
	),
)
