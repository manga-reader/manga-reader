package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserLoginRes struct {
	Token string `json:"token,omitempty"`
}

type JWTTokenClaims struct {
	jwt.StandardClaims
	UserID string `json:"id,omitempty"`
	// ExpiredAt time.Time `json:"expired_at,omitempty"`
}

func UserLogin(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTTokenClaims{
		UserID: c.Param("id"),
	})

	var hmacSampleSecret []byte
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		panic(err)
	}

	res := &UserLoginRes{
		Token: tokenString,
	}

	c.JSON(http.StatusOK, res)
}
