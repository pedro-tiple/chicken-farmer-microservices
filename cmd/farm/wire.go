//go:build wireinject
// +build wireinject

package main

import (
	"chicken-farmer/backend/internal/farm"
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"database/sql"

	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func initializeService(
	address string,
	logger *zap.SugaredLogger,
	dbConnection *sql.DB,
	farmerGRPCConn grpc.ClientConnInterface,
) (farm.Service, error) {
	panic(wire.Build(
		farm.ProvideService,

		farm.ProvideController,
		wire.Bind(new(farm.IController), new(*farm.Controller)),

		farm.ProvideSQLDatabase,
		wire.Bind(new(farm.IDataSource), new(*farm.SQLDatabase)),

		internalGrpc.NewFarmerServiceClient,
		farm.ProvideFarmerService,
		wire.Bind(new(farm.IFarmerService), new(*farm.FarmerService)),
	))
}
