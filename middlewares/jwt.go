package middlewares

import (
	"github.com/gofiber/fiber/v2"
	JWTToken "github.com/gofiber/jwt/v3"
)

// Middleware JWT function
func JWTMiddleware(secret string) fiber.Handler {
	return JWTToken.New(JWTToken.Config{
		SigningKey: []byte(secret),
	})
}
