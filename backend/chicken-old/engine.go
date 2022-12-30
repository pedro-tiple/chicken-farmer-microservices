package chicken_old

import (
	"context"

	"go.uber.org/zap"
)

type Engine struct {
	Database     any // TODO
	MessageQueue any // TODO
	Logger       *zap.SugaredLogger
}

func NewEngine() API {
	return &Engine{}
}

func (engine Engine) GetChicken(ctx context.Context, chickenID string) (Chicken, error) {
	//TODO implement me
	panic("implement me")
}

func (engine Engine) NewChicken(ctx context.Context, farmID string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (engine Engine) FeedChicken(ctx context.Context, chickenID string) error {
	//TODO implement me
	panic("implement me")
}
