package main

import (
	"chicken-farmer/backend/internal/farmer"
	internalGrpc "chicken-farmer/backend/internal/pkg/grpc"
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	farmerService, err := initializeGRPCService(
		ctx, *grpcAddr, logger.Sugar(), farmGRPCConn,
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	log.Println("Service listening")

	go farmerService.ListenForConnections(ctx, farmer.Authenticate)
	go internalGrpc.RunRESTGateway(
		ctx, logger.Sugar(),
		internalGrpc.RegisterFarmerServiceHandlerFromEndpoint,
		*restAddr, *grpcAddr,
	)

	// Wait for termination signal.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	farmerService.GracefulStop()
}
