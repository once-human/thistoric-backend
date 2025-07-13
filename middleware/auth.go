package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"thistoric-backend/core/auth"
)

// JWTSecret loads from environment variables
func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default-secret-key" // Fallback for development
	}
	return secret
}

// AuthMiddleware validates JWT token and sets user info in context
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "🚫 Missing authorization header")
		}

		// Extract token from "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return fiber.NewError(fiber.StatusUnauthorized, "🚫 Invalid authorization header format")
		}

		tokenString := tokenParts[1]
		claims, err := auth.ValidateToken(tokenString, getJWTSecret())
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "🚫 Invalid token")
		}

		// Set user info in context for later use
		c.Locals("user_id", claims.UserID)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}

// RequireRoles checks if the authenticated user has one of the allowed roles
func RequireRoles(allowed ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)
		for _, r := range allowed {
			if role == r {
				return c.Next()
			}
		}
		return fiber.NewError(fiber.StatusForbidden, "🚫 Forbidden: Insufficient role")
	}
}

// RequireAuth is a convenience middleware that combines AuthMiddleware and RequireRoles
func RequireAuth(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// First authenticate
		if err := AuthMiddleware()(c); err != nil {
			return err
		}

		// Then check roles if specified
		if len(allowedRoles) > 0 {
			if err := RequireRoles(allowedRoles...)(c); err != nil {
				return err
			}
		}

		return c.Next()
	}
}
