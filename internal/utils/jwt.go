package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int, jwtSecret string) (string, error) {

	now := time.Now()

	claims := JWTClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string, jwtSecret string) (claims JWTClaims, err error) {
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return claims, err
	}

	return claims, err
}
