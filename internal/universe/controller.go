package universe

import (
	"chicken-farmer/backend/internal/pkg/event"
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

const ()

var ()

type IDataSource interface {
	SetCurrentDay(ctx context.Context, day uint) error
	IncrementCurrentDay(ctx context.Context) (day uint, err error)
}

type Controller struct {
	datasource IDataSource
	publisher  message.Publisher
}

var _ IController = &Controller{}

func ProvideController(
	datasource IDataSource,
	publisher message.Publisher,
) *Controller {
	return &Controller{
		datasource: datasource,
		publisher:  publisher,
	}
}

func (c *Controller) ResetTime(ctx context.Context) error {
	return c.datasource.SetCurrentDay(ctx, 0)
}

func (c *Controller) Tick(ctx context.Context) error {
	newDay, err := c.datasource.IncrementCurrentDay(ctx)
	if err != nil {
		return err
	}

	dayMessageBody, err := json.Marshal(event.DayMessage{Day: newDay})
	if err != nil {
		return err
	}

	msg := message.NewMessage(uuid.New().String(), dayMessageBody)
	msg.Metadata[event.MetadataFieldType] = event.MessageTypeDay

	if err := c.publisher.Publish(event.UniverseTopic, msg); err != nil {
		return err
	}

	return nil
}
