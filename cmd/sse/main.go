package main

import (
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

	restAddr := flag.String("restAddr", "localhost:8081", "REST server address")
	flag.Parse()

	logger, _ := zap.NewProduction()

	// Load environment variables.
	if err := godotenv.Load(); err != nil {
		logger.Fatal(err.Error())
	}

	// TODO from .env
	// Controller will read all events through RabbitMQ.
	amqpConfig := amqp.NewDurablePubSubConfig(
		"amqp://guest:guest@localhost:5672/",
		amqp.GenerateQueueNameTopicNameWithSuffix(uuid.New().String()),
	)
	rabbitMQSubscriber, err := amqp.NewSubscriber(
		amqpConfig, watermill.NewStdLogger(false, false),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sseService, err := initializeHTTPService(
		ctx, logger.Sugar(), rabbitMQSubscriber,
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	log.Println("Service listening")

	go func() {
		if err := sseService.ListenAndServe(ctx, *restAddr); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	// Wait for termination signal.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
}
