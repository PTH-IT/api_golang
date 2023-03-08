package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

func GenerateToken(UserID string) string {
	claims := &MyCustomClaims{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),

			IssuedAt: GettimeNumber(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("test"))
	if err != nil {
		panic(err)
	}
	return t
}
func ParseToken(tokenString string) *jwt.Token {
	token, _ := jwt.Parse(tokenString, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte("test"), nil
	})
	return token
}
