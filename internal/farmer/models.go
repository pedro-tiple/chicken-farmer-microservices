package farmer

import "github.com/google/uuid"

type Farmer struct {
	ID           uuid.UUID
	FarmID       uuid.UUID
	Name         string
	PasswordHash string
}
