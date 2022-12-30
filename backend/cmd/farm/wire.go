//go:build wireinject
// +build wireinject

package main

import (
	"chicken-farmer/backend/farm"
	"database/sql"

	"github.com/google/wire"
	"go.uber.org/zap"
)

func initializeService(
	address string,
	logger *zap.SugaredLogger,
	dbConnection *sql.DB,
) (farm.Service, error) {
	panic(wire.Build(
		farm.ProvideService,

		farm.ProvideController,
		wire.Bind(new(farm.IController), new(*farm.Controller)),

		farm.ProvideSQLDatabase,
		wire.Bind(new(farm.IDatabase), new(*farm.SQLDatabase)),

		farm.ProvideFarmerService,
		wire.Bind(new(farm.IFarmer), new(*farm.FarmerService)),
	))
}
