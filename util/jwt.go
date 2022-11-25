package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

var secretKey = []byte("Helloworld")

func GenerateToken(username string, exptime time.Duration) (string, error) {
	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: exptime.Microseconds(),
			Issuer:    "its me",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(secretKey)
	return token, err
}

//func ParseToken(token string) (*Claims, error) {
//	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return secretKey, nil
//	})
//
//	if tokenClaims != nil {
//		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
//			return claims, nil
//		}
//	}
//	return nil, err
//}
