package jwt

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const (
	// TODO load this password from a .env instead of hardcoded const
	password = "JGJI7ZAc47gEharPnJsQldo2I8Jo4zmfaR8N9enzhgw"

	authorizationHeader = "Authorization"
	signingMethod       = "HS256"
)

// UserClaims holds the claims used for the project's authentication JWT.
// There is no way to revoke these JWT on logout, but for this project that
// isn't important.
type UserClaims struct {
	jwt.StandardClaims

	FarmerID uuid.UUID
}

// ServiceClaims holds the claims used for service authentication.
type ServiceClaims struct {
	jwt.StandardClaims

	FarmerID uuid.UUID
	Service  string
}

// Validate returns an http handler functions that validates a JWT token passed
// on the Authorization header.
func Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Ignore login and ws and OPTIONS requests
			// if (len(r.URL.Path) >= 6 && r.URL.Path[0:6] == "/login") ||
			// 	r.URL.Path == "/ws" ||
			// 	r.Method == http.MethodOptions {
			// 	next.ServeHTTP(w, r)
			// 	return
			// }

			authorization := r.Header.Get(authorizationHeader)
			if authorization == "" {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			authorizationTokens := strings.Split(authorization, " ")
			if len(authorizationTokens) != 2 {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			tokenStr := authorizationTokens[1]
			tk := &UserClaims{}

			token, err := jwt.ParseWithClaims(
				tokenStr, tk, func(token *jwt.Token) (interface{}, error) {
					return []byte(password), nil
				},
			)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			if !token.Valid {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			// TODO const
			ctx := context.WithValue(r.Context(), "jwtToken", tk)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		},
	)
}

func GenerateUserToken(farmerID uuid.UUID) (string, error) {
	return jwt.NewWithClaims(
		jwt.GetSigningMethod(signingMethod), &UserClaims{
			FarmerID: farmerID,
		},
	).SignedString([]byte(password))
}
