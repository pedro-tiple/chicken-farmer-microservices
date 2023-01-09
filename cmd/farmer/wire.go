//go:build wireinject
// +build wireinject

package main

import (
	"chicken-farmer/backend/internal/farmer"
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func initializeService(
	ctx context.Context,
	address string,
	logger *zap.SugaredLogger,
	farmGRPCConn grpc.ClientConnInterface,
) (farmer.Service, error) {
	panic(wire.Build(
		farmer.ProvideService,

		farmer.ProvideController,
		wire.Bind(new(farmer.IController), new(*farmer.Controller)),

		farmer.ProvideMongoDatabase,
		wire.Bind(new(farmer.IDataSource), new(*farmer.MongoDatabase)),

		internalGrpc.NewFarmServiceClient,
		farmer.ProvideFarmService,
		wire.Bind(new(farmer.IFarmService), new(*farmer.FarmService)),
	))
}
