//go:build wireinject
// +build wireinject

package main

import (
	"chicken-farmer/backend/internal/farm"
	farmSql "chicken-farmer/backend/internal/farm/sql"
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"
	"database/sql"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func initializeGRPCService(
	ctx context.Context,
	address string,
	logger *zap.SugaredLogger,
	dbConnection *sql.DB,
	farmerGRPCConn grpc.ClientConnInterface,
	subscriber message.Subscriber,
	publisher message.Publisher,
	jwtAuthKey []byte,
) (*farm.GRPCService, error) {
	panic(
		wire.Build(
			farm.ProvideGRPCService,

			farm.ProvideController,
			wire.Bind(new(farm.IController), new(*farm.Controller)),

			farmSql.ProvideDatasource,
			wire.Bind(new(farm.IDataSource), new(*farmSql.Datasource)),

			internalGrpc.NewFarmerPrivateServiceClient,
			farm.ProvideFarmerPrivateGRPCClient,
			wire.Bind(new(farm.IFarmerService), new(*farm.FarmerGRPCClient)),
		),
	)
}
