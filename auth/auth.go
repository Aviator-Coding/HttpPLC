package auth

import (
	"errors"
	"time"

	"github.com/Aviator-Coding/HttpPLC/configs"
	"github.com/Aviator-Coding/HttpPLC/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Custom defined Claims
// I like to use this but i coudn;t figured out how
type AuthJWTClaims struct {
	jwt.StandardClaims
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

// Middleware to protect the Routes with JWT Token
func AuthHandler(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	// Parse Cookie
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.CFG.JWT.KeySecret), nil
	})

	if err != nil {
		// Return status 401 and failed authentication error.
		if err.Error() == "Missing or malformed JWT" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		// Return status 401 and failed authentication error.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	claims := token.Claims.(jwt.MapClaims)
	c.Locals("name", claims["name"])
	c.Locals("id", claims["id"])
	c.Locals("is_admin", claims["is_admin"])
	c.Locals("email", claims["email"])
	return c.Next()
}

func LoginUser(c *fiber.Ctx, user models.User) (string, error) {

	//Create the Claims
	claims := jwt.MapClaims{
		"name":     user.Name,
		"id":       user.ID,
		"is_admin": user.IsAdmin,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Minute * time.Duration(configs.CFG.JWT.KeyExpireMinutes)).Unix(),
	}

	// Create unsinged token
	tokenUnsigned := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our Secret
	token, err := tokenUnsigned.SignedString([]byte(configs.CFG.JWT.KeySecret))
	if err != nil {
		return "", errors.New("[JWT] - There was an error during token Signature")
	}

	// Create cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return token, nil
}

// Logs out the current User
func LogOutUser(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return nil
}
