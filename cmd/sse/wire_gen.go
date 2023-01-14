// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"chicken-farmer/backend/internal/sse"
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func initializeHTTPService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	sseSubscriber message.Subscriber,
	controllerSubscriber message.Subscriber,
	controllerPublisher message.Publisher,
) (*sse.HTTPService, error) {
	controller, err := sse.ProvideController(ctx, logger, controllerSubscriber, controllerPublisher)
	if err != nil {
		return nil, err
	}
	httpService, err := sse.ProvideWatermillService(logger, controller, sseSubscriber)
	if err != nil {
		return nil, err
	}
	return httpService, nil
}
