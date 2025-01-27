package middleware

import (
	"be-golang/config"
	"be-golang/utils"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var cfg *config.Config

func init() {
	cfg, _ = config.NewConfig()
}

func contains(allowedRoles []string, role interface{}) bool {
	for _, allowedRole := range allowedRoles {
		if role == allowedRole {
			return true
		}
	}
	return false
}

// AuthMiddleware checks the JWT and validates the user's role
func AuthMiddleware(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the token from the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.ErrorResponseHandler(c, fiber.StatusUnauthorized, "Error", "Missing or invalid token")
		}

		// Extract the token (e.g., "Bearer <token>")
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return utils.ErrorResponseHandler(c, fiber.StatusUnauthorized, "Error", "Invalid token format")
		}
		tokenString := tokenParts[1]

		// Parse the token
		secret := cfg.JWT_KEY_SECRET
		if secret == "" {
			return utils.ErrorResponseHandler(c, fiber.StatusInternalServerError, "Error", "Server configuration error")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		// Validate token and claims
		if err != nil || !token.Valid {
			return utils.ErrorResponseHandler(c, fiber.StatusUnauthorized, "Error", "Invalid or expired token")
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return utils.ErrorResponseHandler(c, fiber.StatusUnauthorized, "Error", "Invalid token claims")
		}

		// Extract and validate user data
		userData, ok := claims["data"].(map[string]interface{})
		if !ok || userData == nil {
			return utils.ErrorResponseHandler(c, fiber.StatusUnauthorized, "Error", "Invalid or missing user data in token")
		}

		// Store user data in the Fiber context
		c.Locals("userData", userData)

		// Extract role from user data
		role, ok := userData["role"]
		if !ok {
			return utils.ErrorResponseHandler(c, fiber.StatusUnauthorized, "Error", "Role not found in token")
		}

		// Check if the role is allowed
		if len(allowedRoles) == 0 || contains(allowedRoles, role) {
			return c.Next()
		}

		// Role is not allowed
		return utils.ErrorResponseHandler(c, fiber.StatusForbidden, "Error", "You do not have permission to access this resource")
	}
}
