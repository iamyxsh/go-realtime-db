package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/iamyxsh/go-realtime-db/constants"
)

type JWTClaims struct {
	jwt.StandardClaims
	Email      string `json:"email"`
	Expiration int64  `json:"exp"`
}

func GenerateJWT(email string, expirationTime time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = expirationTime.Unix()

	tokenString, err := token.SignedString([]byte(constants.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("token is invalid")
}
