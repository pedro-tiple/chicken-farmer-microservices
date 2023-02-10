package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const (
	SigningMethod = "HS256"

	ContextKey = "jwt"
)

var (
	ErrInvalidJWT = errors.New("invalid JWT")
)

// UserClaims holds the claims used for the project's authentication JWT.
// There is no way to revoke these JWT on logout, but for this project that
// isn't important.
type UserClaims struct {
	jwt.StandardClaims

	FarmerID uuid.UUID
	FarmID   uuid.UUID
}

// ServiceClaims holds the claims used for service authentication.
type ServiceClaims struct {
	jwt.StandardClaims

	FarmerID uuid.UUID
	Service  string
}

func ValidateUserClaims(key []byte, tokenStr string) (*UserClaims, error) {
	tk := &UserClaims{}
	token, err := jwt.ParseWithClaims(
		tokenStr, tk, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		},
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidJWT
	}

	return tk, nil
}
