// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"chicken-farmer/backend/internal/farm"
	sql2 "chicken-farmer/backend/internal/farm/sql"
	grpc2 "chicken-farmer/backend/internal/pkg/grpc"
	"context"
	"database/sql"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

func initializeGRPCService(ctx context.Context, address string, logger *zap.SugaredLogger, dbConnection *sql.DB, farmerGRPCConn grpc.ClientConnInterface, subscriber message.Subscriber, publisher message.Publisher, jwtAuthKey []byte) (*farm.GRPCService, error) {
	datasource, err := sql2.ProvideDatasource(dbConnection)
	if err != nil {
		return nil, err
	}
	farmerPrivateServiceClient := grpc2.NewFarmerPrivateServiceClient(farmerGRPCConn)
	farmerGRPCClient := farm.ProvideFarmerPrivateGRPCClient(farmerPrivateServiceClient)
	controller, err := farm.ProvideController(ctx, logger, datasource, farmerGRPCClient, subscriber, publisher)
	if err != nil {
		return nil, err
	}
	grpcService := farm.ProvideGRPCService(address, logger, controller, jwtAuthKey)
	return grpcService, nil
}
