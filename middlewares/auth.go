package middleware

import (
	"github.com/fingerprint/constants"
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/services"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	authService      services.AuthService
	organizationRepo repositories.OrganizationRepository
}

func NewAuthMiddleware(authService services.AuthService, organizationRepo repositories.OrganizationRepository) *AuthMiddleware {
	return &AuthMiddleware{
		authService:      authService,
		organizationRepo: organizationRepo,
	}
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

// Organization Guard Middleware

type OrganizationGuardOptions struct {
	OrganizationID string `json:"organization_id"`
}

func (a *AuthMiddleware) OrganizationGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// User
		user := c.Locals("user").(*models.User)

		// Let superadmin pass
		if user.Role == constants.SUPERADMIN.String() {
			return c.Next()
		}

		// Get organization from request
		req := new(OrganizationGuardOptions)
		if err := c.BodyParser(req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		orgs, err := a.organizationRepo.Find(&models.OrganizationFind{ID: req.OrganizationID})
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		if len(*orgs) != 1 {
			return fiber.NewError(fiber.StatusNotFound, "Organization not found")
		}
		org := (*orgs)[0]

		// fmt.Println("User.OrganizationID", user.OrganizationID)
		// fmt.Println("org.ID.String()", org.ID.String())

		// Validation
		if c.Method() == "PUT" {
			if user.OrganizationID != org.ID.String() {
				return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized access to organization")
			}
		}
		return c.Next()
	}
}
