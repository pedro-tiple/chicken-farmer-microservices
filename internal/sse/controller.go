package sse

import (
	"chicken-farmer/backend/internal/pkg/event"
	"context"
	"encoding/json"
	"sync"

	messagePkg "github.com/ThreeDotsLabs/watermill/message"
	"github.com/gin-contrib/sse"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Controller struct {
	logger             *zap.SugaredLogger
	subscriber         messagePkg.Subscriber
	openConnectionsMap map[string][]openConnection
	mu                 sync.Mutex
}

var _ IController = &Controller{}

type openConnection struct {
	ctx     context.Context
	channel chan sse.Event
}

func ProvideController(
	ctx context.Context,
	logger *zap.SugaredLogger,
	subscriber messagePkg.Subscriber,
) (*Controller, error) {
	controller := Controller{
		logger:             logger,
		subscriber:         subscriber,
		openConnectionsMap: make(map[string][]openConnection),
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

	for {
		select {
		case farmMessage := <-farmChan:
			if err := c.processFarmMessage(farmMessage); err != nil {
				c.logger.Error(err.Error())
				continue
			}
		case universeMessage := <-universeChan:
			if err := c.processUniverseMessage(universeMessage); err != nil {
				c.logger.Error(err.Error())
				continue
			}
		}
	}
}

func (c *Controller) processFarmMessage(message *messagePkg.Message) error {
	farmerID := message.Metadata[event.MetadataFieldFarmerID]
	if openConnections, ok := c.openConnectionsMap[farmerID]; ok {
		c.sendEventToConnections(
			sse.Event{
				Event: message.Metadata[event.MetadataFieldType],
				Data:  string(message.Payload),
			}, &openConnections,
		)
	}

	message.Ack()

	return nil
}

func (c *Controller) processUniverseMessage(message *messagePkg.Message) error {
	var dayMessage event.DayMessage
	if err := json.Unmarshal(message.Payload, &dayMessage); err != nil {
		return err
	}

	sseEvent := sse.Event{
		Event: event.MessageTypeDay,
		Data:  dayMessage.Day,
	}

	for _, openConnections := range c.openConnectionsMap {
		c.sendEventToConnections(sseEvent, &openConnections)
	}

	message.Ack()

	return nil
}

func (c *Controller) sendEventToConnections(
	event sse.Event, openConnections *[]openConnection,
) {
	for k, oc := range *openConnections {
		if oc.channel == nil {
			continue
		}

		// Remove closed connections
		if oc.ctx.Err() != nil {
			// TODO here we would want to delete the entry in openConnections
			// but removing while ranging causes issues and couldn't find a
			// simple alternative.
			close(oc.channel)
			(*openConnections)[k] = openConnection{}
			continue
		}

		oc.channel <- event
	}
}

func (c *Controller) SubscribeToFarmer(
	ctx context.Context, farmerID uuid.UUID,
) (chan sse.Event, error) {
	newOC := openConnection{ctx: ctx, channel: make(chan sse.Event)}
	c.mu.Lock()
	defer c.mu.Unlock()
	if oc, ok := c.openConnectionsMap[farmerID.String()]; ok {
		c.openConnectionsMap[farmerID.String()] = append(oc, newOC)
	} else {
		c.openConnectionsMap[farmerID.String()] = []openConnection{newOC}
	}

	return newOC.channel, nil
}
