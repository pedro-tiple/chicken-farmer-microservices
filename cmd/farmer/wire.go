//go:build wireinject
// +build wireinject

package main

import (
	"chicken-farmer/backend/internal/farmer"
	farmerMongo "chicken-farmer/backend/internal/farmer/mongo"
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func initializeGRPCService(
	ctx context.Context,
	address string,
	logger *zap.SugaredLogger,
	farmGRPCConn grpc.ClientConnInterface,
) (*farmer.GRPCService, error) {
	panic(
		wire.Build(
			farmer.ProvideGRPCService,

			farmer.ProvideController,
			wire.Bind(new(farmer.IController), new(*farmer.Controller)),

			farmerMongo.ProvideDatasource,
			wire.Bind(new(farmer.IDataSource), new(*farmerMongo.Datasource)),

			internalGrpc.NewFarmServiceClient,
			farmer.ProvideFarmGRPCClient,
			wire.Bind(new(farmer.IFarmService), new(*farmer.FarmGRPCClient)),
		),
	)
}
