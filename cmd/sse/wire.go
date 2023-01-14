//go:build wireinject
// +build wireinject

package main

import (
	"chicken-farmer/backend/internal/sse"
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func initializeHTTPService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	sseSubscriber message.Subscriber,
	controllerSubscriber message.Subscriber,
	controllerPublisher message.Publisher,
) (*sse.HTTPService, error) {
	panic(
		wire.Build(
			sse.ProvideWatermillService,

			sse.ProvideController,
			wire.Bind(new(sse.IController), new(*sse.Controller)),
		),
	)
}
