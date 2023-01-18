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
	subscriber message.Subscriber,
) (*sse.HTTPService, error) {
	panic(
		wire.Build(
			sse.ProvideHTTPService,

			sse.ProvideController,
			wire.Bind(new(sse.IController), new(*sse.Controller)),
		),
	)
}
