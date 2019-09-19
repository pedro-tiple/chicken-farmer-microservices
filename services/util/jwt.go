package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

type JwtToken struct {
	UserId    primitive.ObjectID
	IsService bool
	Token     string
	jwt.StandardClaims
}

// TODO load this password from a .env instead of hardcoded const
const password = "JGJI7ZAc47gEharPnJsQldo2I8Jo4zmfaR8N9enzhgw"

// There is no way to revoke these jwt tokens on logout, but for this project that isn't important
func JwtAuthentication(_next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ignore login and ws and OPTIONS requests
		if (len(r.URL.Path) >= 6 && r.URL.Path[0:6] == "/login") ||
			r.URL.Path == "/ws" ||
			r.Method == http.MethodOptions {
			_next.ServeHTTP(w, r)
			return
		}

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		strTokens := strings.Split(tokenHeader, " ")
		if len(strTokens) != 2 {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		token := strTokens[1]
		tk := &JwtToken{Token: token}

		jwtToken, err := jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(password), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if !jwtToken.Valid {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "jwtToken", tk)
		r = r.WithContext(ctx)
		_next.ServeHTTP(w, r)
	})
}

func GenerateJwtToken(_userId primitive.ObjectID, _isService bool) (string, error) {
	tk := &JwtToken{
		UserId:    _userId,
		IsService: _isService,
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	return token.SignedString([]byte(password))
}

func BuildRequest(_requestType string, _url string, _body interface{}, _from primitive.ObjectID) (*http.Request, error) {
	token, err := GenerateJwtToken(_from, true)
	if err != nil {
		return nil, err
	}

	requestBody, err := json.Marshal(_body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(_requestType, _url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return request, nil
}
