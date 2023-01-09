//go:build wireinject
// +build wireinject

package main

import (
	"chicken-farmer/backend/internal/farmer"
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
)

func initializeService(
	ctx context.Context,
	address string,
	logger *zap.SugaredLogger,
) (farmer.Service, error) {
	panic(wire.Build(
		farmer.ProvideService,

		farmer.ProvideController,
		wire.Bind(new(farmer.IController), new(*farmer.Controller)),

		farmer.ProvideMongoDatabase,
		wire.Bind(new(farmer.IDataSource), new(*farmer.MongoDatabase)),
	))
}
