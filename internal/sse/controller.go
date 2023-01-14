package sse

import (
	"chicken-farmer/backend/internal/pkg/event"
	"context"
	"encoding/json"
	"fmt"

	messagePkg "github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Controller struct {
	logger            *zap.SugaredLogger
	subscriber        messagePkg.Subscriber
	publisher         messagePkg.Publisher
	registeredFarmers map[string]context.Context
}

var _ IController = &Controller{}

func ProvideController(
	ctx context.Context,
	logger *zap.SugaredLogger,
	subscriber messagePkg.Subscriber,
	publisher messagePkg.Publisher,
) (*Controller, error) {
	controller := Controller{
		logger:            logger,
		subscriber:        subscriber,
		publisher:         publisher,
		registeredFarmers: make(map[string]context.Context),
	}

	go func() {
		if err := controller.processMessages(ctx); err != nil {
			// TODO recover
			controller.logger.Error(err.Error())
		}
	}()

	return &controller, nil
}

func (c *Controller) processMessages(ctx context.Context) error {
	universeChan, err := c.subscriber.Subscribe(ctx, event.UniverseTopic)
	if err != nil {
		return err
	}

	farmChan, err := c.subscriber.Subscribe(ctx, event.FarmTopic)
	if err != nil {
		return err
	}

	var (
		topic   string
		message *messagePkg.Message
	)
OuterLoop:
	for {
		select {
		case farmMessage := <-farmChan:
			farmerID := farmMessage.Metadata[event.MetadataFieldFarmerID]
			if _, ok := c.registeredFarmers[farmerID]; ok {
				message, err = farmMessageToEventMessage(
					string(farmMessage.Payload),
				)
				if err != nil {
					c.logger.Error(err.Error())
					continue
				}

				topic = fmt.Sprintf(event.UserEventsTopic, farmerID)
				farmMessage.Ack()
				// Continue to single message publish.
			}
		case universeMessage := <-universeChan:
			// Single universe message is fanned out to all farmers.
			for k, v := range c.registeredFarmers {
				// Remove closed connections.
				if v.Err() != nil {
					delete(c.registeredFarmers, k)
					continue
				}

				message, err = universeMessageToEventMessage(
					string(universeMessage.Payload),
				)
				if err != nil {
					c.logger.Error(err.Error())
					continue
				}

				if err = c.publisher.Publish(
					fmt.Sprintf(event.UserEventsTopic, k), message,
				); err != nil {
					c.logger.Error(err.Error())
					continue
				}
			}

			universeMessage.Ack()

			// Don't continue to single message publish.
			continue OuterLoop
		}

		if err := c.publisher.Publish(topic, message); err != nil {
			c.logger.Error(err.Error())
			continue
		}
	}
}

func farmMessageToEventMessage(message string) (*messagePkg.Message, error) {
	// TODO convert event to change

	messagePayload, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	return messagePkg.NewMessage(uuid.NewString(), messagePayload), nil
}

func universeMessageToEventMessage(message string) (*messagePkg.Message, error) {
	fmt.Println(message)
	return messagePkg.NewMessage(uuid.NewString(), []byte(message)), nil
}

func (c Controller) RegisterFarmer(
	ctx context.Context, farmerID uuid.UUID,
) error {
	c.registeredFarmers[farmerID.String()] = ctx

	return nil
}
