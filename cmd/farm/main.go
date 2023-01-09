package main

import (
	"chicken-farmer/backend/internal/farm"
	internalDB "chicken-farmer/backend/internal/pkg/database"
	cfGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	log.Println("Setting up things...")
	logger, _ := zap.NewProduction()

	grpcAddr := flag.String("grpcAddr", "localhost:50051", "gRPC server address")
	restAddr := flag.String("restAddr", "localhost:8081", "REST server address")
	flag.Parse()

	// Load environment variables.
	if err := godotenv.Load(); err != nil {
		logger.Fatal(err.Error())
	}

	dbConnections, err := internalDB.OpenSQLConnections([]internalDB.ConnectionSettings{{
		Host:          os.Getenv("POSTGRES_HOST"),
		Port:          os.Getenv("POSTGRES_PORT"),
		DatabaseName:  os.Getenv("POSTGRES_DB"),
		User:          os.Getenv("POSTGRES_USER"),
		Password:      os.Getenv("POSTGRES_PASSWORD"),
		IsReadReplica: false,
	}})
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer func() {
		for _, connection := range dbConnections {
			_ = connection.Close()
		}
	}()

	driver, err := postgres.WithInstance(dbConnections[0], &postgres.Config{})
	if err != nil {
		logger.Fatal(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		os.Getenv("MIGRATIONS_FOLDER"), os.Getenv("POSTGRES_DB"), driver,
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	//if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
	//	logger.Fatal(err.Error())
	//}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatal(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	farmerGRPCConn, err := cfGrpc.CreateClientConnection(
		ctx, os.Getenv("FARMER_SERVICE_ADDRESS"),
	)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer farmerGRPCConn.Close()

	farmService, err := initializeService(
		*grpcAddr,
		logger.Sugar(),
		dbConnections[0],
		farmerGRPCConn,
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	log.Println("Service listening")

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	go farmService.ListenForConnections(ctx, farm.Authenticate)
	go cfGrpc.RunRESTGateway(
		ctx, logger.Sugar(),
		cfGrpc.RegisterFarmServiceHandlerFromEndpoint,
		*restAddr, *grpcAddr,
	)

	// Wait for termination signal.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	farmService.GracefulStop()
}
