package chicken_old

import "context"

// To generate mocks:
// mockgen -package proto-old -source interface.go -destination mock.go

type API interface {
	// GetChicken returns the stored Chicken associated to the given ID.
	GetChicken(ctx context.Context, chickenID string) (Chicken, error)

	// NewChicken creates a new Chicken assigned to the provided proto.
	// Returns the ID of the newly created proto-old.
	NewChicken(ctx context.Context, farmID string) (string, error)

	// FeedChicken makes the proto-old lay an egg.
	// Each proto-old can only be fed once per universe tick.
	FeedChicken(ctx context.Context, chickenID string) error
}
