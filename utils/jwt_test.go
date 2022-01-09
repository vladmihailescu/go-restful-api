package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJwt(t *testing.T) {
	id := 7
	jwt, err := GenerateToken(uint(id))
	require.Nil(t, err)
	require.NotEqual(t, jwt, id)

	idFromJwt, err := ValidateToken(jwt)
	require.Nil(t, err)
	require.Equal(t, int(idFromJwt), id)
}

func TestHashPassword(t *testing.T) {
	password := "asd"
	hashPassword := HashPassword(password)

	require.NotEqual(t, password, hashPassword)
	require.Equal(t, hashPassword, HashPassword("asd"))
}
