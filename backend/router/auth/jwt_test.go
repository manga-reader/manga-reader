package auth_test

import (
	"testing"

	"github.com/manga-reader/manga-reader/backend/router/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_JWT_EncodeDecode(t *testing.T) {
	userID := "john"

	jwtStr, err := auth.EncodeJWT(userID)
	require.NoError(t, err)
	claim, err := auth.DecodeJWT(jwtStr)
	assert.NoError(t, err)
	assert.Equal(t, userID, claim.UserID)
}
