package utils

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	// hmacSampleSecret must be specified
	hmacSampleSecret []byte

	oneDay = 24 * time.Hour

	bitsForInt64 = 64

	claimBase = 10
)

func GenerateToken(id uint) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(oneDay).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", fmt.Errorf("unable to generate token: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (uint, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	if err != nil {
		return 0, fmt.Errorf("unable to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("unable to get claims from token")
	}
	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	if time.Now().Unix() > int64(claims["exp"].(float64)) {
		return 0, fmt.Errorf("token expired")
	}

	return uint(claims["id"].(float64)), nil
}

func HashPassword(password string) string {
	bytePassword := sha256.Sum256([]byte(password))
	password = string(bytePassword[:])
	return password
}
