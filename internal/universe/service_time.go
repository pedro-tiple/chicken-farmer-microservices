package universe

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type IController interface {
	ResetTime(ctx context.Context) error
	Tick(ctx context.Context) error
}

type Service struct {
	logger     *zap.SugaredLogger
	controller IController
	frequency  time.Duration
}

func ProvideService(
	logger *zap.SugaredLogger,
	controller IController,
	frequency time.Duration,
) Service {
	return Service{
		logger:     logger,
		controller: controller,
		frequency:  frequency,
	}
}

func (s Service) BigBang(ctx context.Context) error {
	if err := s.controller.ResetTime(ctx); err != nil {
		return err
	}

	ticker := time.NewTicker(s.frequency)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := s.controller.Tick(ctx); err != nil {
				return err
			}
		}
	}
}
