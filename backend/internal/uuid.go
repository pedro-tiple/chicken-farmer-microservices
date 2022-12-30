package internal

import (
	"math/big"

	"github.com/google/uuid"
)

// UUIDFromString converts a string to UUID ignoring errors.
func UUIDFromString(s string) uuid.UUID {
	result, _ := uuid.Parse(s)
	return result
}

// UUIDFromInt64 converts an int64 to UUID.
// Sqlc returns insert ID as int, so we need to convert to bytes and that's why
// this exists.
func UUIDFromInt64(i int64) (uuid.UUID, error) {
	bigInt := new(big.Int)
	bigInt.SetInt64(i)
	return uuid.FromBytes(bigInt.Bytes())
}
