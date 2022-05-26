package auth

import "github.com/golang-jwt/jwt"

type JWTTokenClaims struct {
	jwt.StandardClaims
	UserID string `json:"id,omitempty"`
}

func GenerateNewToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTTokenClaims{
		UserID: id,
	})

	var hmacSampleSecret []byte
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
