package util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var secret []byte

func init() {
	secret = []byte("SUPER_DUPER_SECRET_KEY")
}

func GenerateToken(email, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issuer": "github.com/erik-olsson-op/go-rest",
		"aud":    "the Internet",
		"email":  email,
		"sub":    userId,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
		"jti":    uuid.New(),
	})
	return token.SignedString(secret)
}

func VerifyToken(token string) (string, string, error) {
	pt, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexcpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return "", "", errors.New("failed to parse token")
	}
	if !pt.Valid {
		return "", "", errors.New("invalid token")
	}
	claims, ok := pt.Claims.(jwt.MapClaims)

	if !ok {
		return "", "", errors.New("invalid token claims")
	}

	email := claims["email"].(string)
	userId := claims["sub"].(string)
	return email, userId, nil
}
