package utils

import (
	"time"

	"github.com/Aviator-Coding/HttpPLC/configs"
	"github.com/golang-jwt/jwt"
)

// GenerateNewAccessToken func for generate a new Access token.
func GenerateNewAccessToken() (string, error) {
	// Set secret key from .env file.
	secret := configs.CFG.JWT.KeySecret

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims with expires minutes count for secret key :
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(configs.CFG.JWT.KeyExpireMinutes)).Unix()

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
