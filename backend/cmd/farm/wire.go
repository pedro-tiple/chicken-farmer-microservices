//go:build wireinject
// +build wireinject

package main

import (
	farm2 "chicken-farmer/backend/internal/farm"
	"database/sql"

	"github.com/google/wire"
	"go.uber.org/zap"
)

func initializeService(
	address string,
	logger *zap.SugaredLogger,
	dbConnection *sql.DB,
) (farm2.Service, error) {
	panic(wire.Build(
		farm2.ProvideService,

		farm2.ProvideController,
		wire.Bind(new(farm2.IController), new(*farm2.Controller)),

		farm2.ProvideSQLDatabase,
		wire.Bind(new(farm2.IDatabase), new(*farm2.SQLDatabase)),

		farm2.ProvideFarmerService,
		wire.Bind(new(farm2.IFarmer), new(*farm2.FarmerService)),
	))
}
