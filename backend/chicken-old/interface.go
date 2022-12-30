package chicken_old

import "context"

// To generate mocks:
// mockgen -package farm-old -source interface.go -destination mock.go

type API interface {
	// GetChicken returns the stored Chicken associated to the given ID.
	GetChicken(ctx context.Context, chickenID string) (Chicken, error)

	// NewChicken creates a new Chicken assigned to the provided farm.
	// Returns the ID of the newly created farm-old.
	NewChicken(ctx context.Context, farmID string) (string, error)

	// FeedChicken makes the farm-old lay an egg.
	// Each farm-old can only be fed once per time tick.
	FeedChicken(ctx context.Context, chickenID string) error
}
