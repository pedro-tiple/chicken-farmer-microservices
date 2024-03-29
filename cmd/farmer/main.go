package main

import (
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
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

	farmGRPCConn, err := internalGrpc.CreateClientConnection(
		context.Background(), os.Getenv("FARM_SERVICE_ADDRESS"),
	)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer farmGRPCConn.Close()

	publisher, err := amqp.NewPublisher(
		amqp.NewDurablePubSubConfig(
			"amqp://guest:guest@localhost:5672/",
			amqp.GenerateQueueNameTopicNameWithSuffix(uuid.New().String()),
		), watermill.NewStdLogger(false, false),
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	farmerService, err := initializeGRPCService(
		ctx, *grpcAddr, logger.Sugar(), farmGRPCConn, publisher,
		[]byte(os.Getenv("JWT_AUTH_KEY")),
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	log.Println("Service listening")

	go farmerService.ListenForConnections(ctx)
	go internalGrpc.RunRESTGateway(
		ctx, logger.Sugar(),
		internalGrpc.RegisterFarmerPublicServiceHandlerFromEndpoint,
		*restAddr, *grpcAddr,
	)

	// Wait for termination signal.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	farmerService.GracefulStop()
}
