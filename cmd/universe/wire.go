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

func initializeTimeService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	frequency time.Duration,
	publisher message.Publisher,
) (*universe.TimeService, error) {
	panic(
		wire.Build(
			universe.ProvideTimeService,

			universe.ProvideController,
			wire.Bind(new(universe.IController), new(*universe.Controller)),

			universe.ProvideMemoryDatasource,
			wire.Bind(new(universe.IDataSource), new(*universe.Datasource)),
		),
	)
}
