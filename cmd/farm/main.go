package main

import (
	"chicken-farmer/backend/internal/farm"
	internalDB "chicken-farmer/backend/internal/pkg/database"
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	// TODO this is too much code and a lot of duplication with other mains, split and clean up.
	log.Println("Setting up things...")

	logger, _ := zap.NewProduction()

	grpcAddr := flag.String(
		"grpcAddr", "localhost:50051", "gRPC server address",
	)
	restAddr := flag.String("restAddr", "localhost:8081", "REST server address")
	flag.Parse()

	// Load environment variables.
	if err := godotenv.Load(); err != nil {
		logger.Fatal(err.Error())
	}

	dbConnections, err := internalDB.OpenSQLConnections(
		[]internalDB.ConnectionSettings{{
			Host:          os.Getenv("POSTGRES_HOST"),
			Port:          os.Getenv("POSTGRES_PORT"),
			DatabaseName:  os.Getenv("POSTGRES_DB"),
			User:          os.Getenv("POSTGRES_USER"),
			Password:      os.Getenv("POSTGRES_PASSWORD"),
			IsReadReplica: false,
		}},
	)
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

	migrator, err := migrate.NewWithDatabaseInstance(
		os.Getenv("MIGRATIONS_FOLDER"), os.Getenv("POSTGRES_DB"), driver,
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// if err := migrator.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
	//	logger.Fatal(err.Error())
	// }

	if err := migrator.Up(); err != nil && !errors.Is(
		err, migrate.ErrNoChange,
	) {
		logger.Fatal(err.Error())
	}

	farmerGRPCConn, err := internalGrpc.CreateClientConnection(
		context.Background(), os.Getenv("FARMER_SERVICE_ADDRESS"),
	)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer farmerGRPCConn.Close()

	// TODO from .env
	amqpConfig := amqp.NewDurablePubSubConfig(
		"amqp://guest:guest@localhost:5672/",
		amqp.GenerateQueueNameTopicNameWithSuffix(uuid.New().String()),
	)
	subscriber, err := amqp.NewSubscriber(
		amqpConfig, watermill.NewStdLogger(false, false),
	)
	if err != nil {
		logger.Fatal(err.Error())
	}
	publisher, err := amqp.NewPublisher(
		amqpConfig, watermill.NewStdLogger(false, false),
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	farmService, err := initializeGRPCService(
		ctx,
		*grpcAddr,
		logger.Sugar(),
		dbConnections[0],
		farmerGRPCConn,
		subscriber,
		publisher,
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	log.Println("Service listening")

	go farmService.ListenForConnections(ctx, farm.Authenticate)
	go internalGrpc.RunRESTGateway(
		ctx, logger.Sugar(),
		internalGrpc.RegisterFarmServiceHandlerFromEndpoint,
		*restAddr, *grpcAddr,
	)

	// Wait for termination signal.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	farmService.GracefulStop()
}
