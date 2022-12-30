package main

import (
	"chicken-farmer/backend/farm"
	internalDB "chicken-farmer/backend/internal/database"
	"context"
	"errors"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"golang.org/x/exp/slog"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	logger, _ := zap.NewProduction()
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr)))
	slog.Info("Setting up things...")

	grpcAddr := flag.String("grpcAddr", "localhost:50051", "gRPC server address")
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
		os.Getenv("MIGRATIONS_FOLDER"), "chicken-farmer", driver,
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatal(err.Error())
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatal(err.Error())
	}

	farmService, err := initializeService(
		*grpcAddr,
		logger.Sugar(),
		dbConnections[0],
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	slog.Info("Service listening")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go farmService.ListenForConnections(ctx, farm.Authenticate)

	// Wait for termination signal.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("Shutting down server...")
	farmService.GracefulStop()
}
