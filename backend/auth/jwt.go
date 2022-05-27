package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/manga-reader/manga-reader/backend/config"
)

type JWTTokenClaims struct {
	jwt.StandardClaims
	UserID string `json:"id,omitempty"`
}

var jwtHMACSecret []byte

func init() {
	jwtHMACSecret = []byte(config.Cfg.Connection.JWTSecret)
}

func GenerateJWTString(id string) (string, error) {
	return EncodeJWT(id)
}

func EncodeJWT(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTTokenClaims{
		UserID: id,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtHMACSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeJWT(tokenString string) (*JWTTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtHMACSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token string: %v: %w", tokenString, err)
	}

	claims, ok := token.Claims.(*JWTTokenClaims)
	if ok && token.Valid {
		goto success
	}
	return nil, fmt.Errorf("can't type assert to JWTTokenClaims or token is not valid with claims: %v", token.Claims)

success:
	return claims, nil
}
