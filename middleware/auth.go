package middleware

import (
	"rizkiwhy-dating-app/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Protected protect routes
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract JWT from the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return jwtError(c, fiber.NewError(fiber.StatusUnauthorized, "Missing or malformed JWT"))
		}

		// Remove "Bearer " prefix
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse the JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid JWT signing method")
			}
			return []byte(config.Config("SECRET")), nil
		})

		if err != nil {
			return jwtError(c, err)
		}

		// Extract claims and set user_id in context
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return jwtError(c, fiber.NewError(fiber.StatusUnauthorized, "Invalid JWT claims"))
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			return jwtError(c, fiber.NewError(fiber.StatusUnauthorized, "user_id not found in JWT"))
		}

		// Set user_id to context
		c.Locals("user_id", userID)

		// Call the next middleware or handler
		return c.Next()
	}
}

func jwtError(c *fiber.Ctx, err error) error {
	var statusCode int
	var message string

	switch err.Error() {
	case "Missing or malformed JWT":
		statusCode = fiber.StatusBadRequest
		message = "Missing or malformed JWT"
	case "Invalid JWT signing method":
		statusCode = fiber.StatusUnauthorized
		message = "Invalid JWT signing method"
	case "Invalid JWT claims":
		statusCode = fiber.StatusUnauthorized
		message = "Invalid JWT claims"
	case "user_id not found in JWT":
		statusCode = fiber.StatusUnauthorized
		message = "user_id not found in JWT"
	default:
		statusCode = fiber.StatusUnauthorized
		message = "Invalid or expired JWT"
	}

	return c.Status(statusCode).JSON(fiber.Map{"status": "error", "message": message, "data": nil})
}
