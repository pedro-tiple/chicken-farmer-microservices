package ctxFarm

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

// Based on https://github.com/grpc-ecosystem/go-grpc-middleware/tree/master/logging/zap/ctxzap

type ctxMarker struct{}

var (
	ctxMarkerKey          = &ctxMarker{}
	ErrMissingFarmContext = errors.New("farm context data is missing")
)

type Data struct {
	FarmerID uuid.UUID
	FarmID   uuid.UUID
}

// SetInContext returns a new context with the provided farm data.
func SetInContext(ctx context.Context, farmerID, farmID uuid.UUID) context.Context {
	return context.WithValue(ctx, ctxMarkerKey, &Data{
		FarmerID: farmerID,
		FarmID:   farmID,
	})
}

// Extract takes the call-scoped farm identification from the context.
func Extract(ctx context.Context) (*Data, error) {
	ctxFarm, ok := ctx.Value(ctxMarkerKey).(*Data)
	if !ok || ctxFarm == nil {
		return nil, ErrMissingFarmContext
	}

	return ctxFarm, nil
}
