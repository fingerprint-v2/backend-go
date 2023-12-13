package middleware

import (
	"github.com/fingerprint/services"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	authService services.AuthService
}

// NewAuthMiddleware creates a new instance of AuthMiddleware.
func NewAuthMiddleware(authService services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{authService: authService}
}

func (a *AuthMiddleware) Auth() fiber.Handler {
	return a.validateJWT()
}

// validateJWT is a middleware function for JWT validation.
func (a *AuthMiddleware) validateJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ignorePathInst, err := utils.NewIgnorePathInstance([]string{
			"/hello-world",
			"/api/v1/ping",
			"/swagger/*",
			"/api/v1/login",
		})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		isIgnorePath := ignorePathInst(c)
		if isIgnorePath {
			return c.Next()
		}

		token := c.Cookies("access_token")

		user, err := a.authService.ValidateToken(token)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		c.Locals("user", user)

		return c.Next()
	}
}
