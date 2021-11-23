package token

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var userId uint64 = 1

func createToken(t *testing.T)  string {
	newToken, err := CreateToken(userId)

	if err != nil {
		t.Errorf(err.Error())
	}
	return newToken
}

func TestCreateToken(t *testing.T) {
	newToken:=createToken(t)

	require.NotEmpty(t, newToken)
}

func TestIsExpiredFalse(t *testing.T) {
	newToken:=createToken(t)

	token, err := ParseToken(newToken)
	require.NoError(t, err)

	ok, err := token.IsExpired()
	require.NoError(t, err)
	require.False(t, ok)
}

func TestUserId(t *testing.T) {
	newToken:=createToken(t)

	token, err := ParseToken(newToken)
	require.NoError(t, err)

	uId, err := token.UserId()
	require.NoError(t, err)
	require.Equal(t, userId, uId)
}
