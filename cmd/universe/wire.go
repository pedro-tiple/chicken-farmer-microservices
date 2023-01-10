//go:build wireinject
// +build wireinject

package main

import (
	"chicken-farmer/backend/internal/universe"
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func initializeService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	frequency time.Duration,
	publisher message.Publisher,
) (universe.Service, error) {
	panic(
		wire.Build(
			universe.ProvideService,

			universe.ProvideController,
			wire.Bind(new(universe.IController), new(*universe.Controller)),

			universe.ProvideMemoryDatasource,
			wire.Bind(new(universe.IDataSource), new(*universe.Datasource)),
		),
	)
}
