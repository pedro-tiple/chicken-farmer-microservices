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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	farmerService, err := initializeService(ctx, *grpcAddr, logger.Sugar())
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
