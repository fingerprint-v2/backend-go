package middleware

import (
	"github.com/fingerprint/constants"
	"github.com/fingerprint/models"
	"github.com/fingerprint/services"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	authService services.AuthService
}

func NewAuthMiddleware(authService services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{authService: authService}
}

// JWT Middleware
func (a *AuthMiddleware) ValidateJWT() fiber.Handler {
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

// Role Guard Middleware

func (a *AuthMiddleware) SuperAdminGuard() fiber.Handler {
	return a.validateUserByRole([]constants.UserRole{constants.SUPERADMIN})
}

func (a *AuthMiddleware) AdminGuard() fiber.Handler {
	return a.validateUserByRole([]constants.UserRole{constants.SUPERADMIN, constants.ADMIN})
}

func (a *AuthMiddleware) validateUserByRole(roles []constants.UserRole) fiber.Handler {
	return func(c *fiber.Ctx) error {

		user := c.Locals("user").(*models.User)
		for _, value := range roles {
			if user.Role == value.String() {
				return c.Next()
			}
		}
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")

	}
}
