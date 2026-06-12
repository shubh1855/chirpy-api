package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// valid jwt
func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()

	token, err := MakeJWT(
		userID,
		"secret",
		time.Hour,
	)
	require.NoError(t, err)

	id, err := ValidateJWT(
		token,
		"secret",
	)

	require.NoError(t, err)
	require.Equal(t, userID, id)
}

// wrong secret
func TestWrongSecret(t *testing.T) {
	userID := uuid.New()

	token, _ := MakeJWT(
		userID,
		"secret",
		time.Hour,
	)

	_, err := ValidateJWT(
		token,
		"wrong-secret",
	)

	require.Error(t, err)
}

// expired token
func TestExpiredJWT(t *testing.T) {
	userID := uuid.New()

	token, _ := MakeJWT(
		userID,
		"secret",
		-time.Hour,
	)

	_, err := ValidateJWT(
		token,
		"secret",
	)

	require.Error(t, err)
}
