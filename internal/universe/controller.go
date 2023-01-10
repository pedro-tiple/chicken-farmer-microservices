package universe

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

const (
	DayTopic = "universe.day.topic"
)

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

	dayMessage, err := json.Marshal(DayMessage{Day: newDay})
	if err != nil {
		return err
	}

	if err := c.publisher.Publish(
		DayTopic, message.NewMessage(uuid.New().String(), dayMessage),
	); err != nil {
		panic(err)
	}

	return nil
}
