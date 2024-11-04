package middleware

import (
	"context"
	"fmt"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	tokenTTL        = 24 * time.Hour
	authHeaderName  = "X-Token"
	contextTokenKey = "user-token"
)

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleCustomer UserRole = "customer"
)

type UserJWT struct {
	Id    int      `json:"id"`
	Role  UserRole `json:"role"`
	Alias string   `json:"alias"`
	jwt.StandardClaims
}

func GenerateToken(id int, role UserRole, alias, jwtSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserJWT{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Id:    id,
		Role:  role,
		Alias: alias,
	})

	return token.SignedString([]byte(jwtSecret))
}

func Auth(jwtSecret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token := r.Header.Get(authHeaderName)

			if token == "" {
				logger.Error(r.Context(), "middleware auth failed: empty token")
				http.Error(w, "Token is required", http.StatusForbidden)
				return
			}

			ctx, err := addTokenToContext(r.Context(), token, jwtSecret)
			if err != nil {
				logger.Error(r.Context(), "middleware auth failed: invalid token")
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func addTokenToContext(ctx context.Context, token, jwtSecret string) (context.Context, error) {
	userJWT, err := parseToken(token, jwtSecret)
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, contextTokenKey, userJWT), nil
}

func parseToken(t, jwtSecret string) (userToken *UserJWT, err error) {
	token, err := jwt.ParseWithClaims(t, &UserJWT{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signin method")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	userToken, ok := token.Claims.(*UserJWT)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return userToken, nil
}
