package middleware

import (
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
