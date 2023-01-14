package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	log.Println("Setting up things...")

	logger, _ := zap.NewProduction()

	// Load environment variables.
	if err := godotenv.Load(); err != nil {
		logger.Fatal(err.Error())
	}

	// TODO amqp string on .env
	// TODO log to zap
	publisher, err := amqp.NewPublisher(
		amqp.NewDurablePubSubConfig(
			"amqp://guest:guest@localhost:5672/",
			amqp.GenerateQueueNameTopicNameWithSuffix(uuid.New().String()),
		),
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	universeService, err := initializeTimeService(
		ctx, logger.Sugar(), time.Second, publisher,
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	log.Println("Service listening")

	go func() {
		if err := universeService.BigBang(ctx); err != nil {
			logger.Error(err.Error())
		}
	}()

	// Wait for termination signal.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
}
