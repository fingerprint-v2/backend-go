package middleware

import (
	"strings"

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
	userRepo         repositories.UserRepository
}

func NewAuthMiddleware(
	authService services.AuthService,
	organizationRepo repositories.OrganizationRepository,
	userRepo repositories.UserRepository,
) *AuthMiddleware {
	return &AuthMiddleware{
		authService:      authService,
		organizationRepo: organizationRepo,
		userRepo:         userRepo,
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
	All            bool   `json:"all"`
	ID             string `json:"id"`
}

func (a *AuthMiddleware) OrganizationGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// User
		user := c.Locals("user").(*models.User)

		// Let superadmin pass
		if user.Role == constants.SUPERADMIN.String() {
			return c.Next()
		}

		var organizationID string
		//
		req := new(OrganizationGuardOptions)
		if err := c.BodyParser(req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		// Non-superadmin cannot set "all" field to true
		if req.All {
			return fiber.NewError(fiber.StatusUnauthorized, "Not allow all field to be true")
		}
		// Make sure that user input organization id if it is not DELETE method
		if req.OrganizationID == "" && c.Method() != "DELETE" {
			return fiber.NewError(fiber.StatusBadRequest, "organization_id is required")
		}
		organizationID = req.OrganizationID

		// If it is PUT, PATCH, or POST method, we need to check if the organization_id is valid
		if c.Method() == "PUT" || c.Method() == "PATCH" || c.Method() == "POST" {
			orgs, err := a.organizationRepo.Find(&models.OrganizationFind{ID: req.OrganizationID})
			if err != nil {
				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}
			if len(*orgs) != 1 {
				return fiber.NewError(fiber.StatusNotFound, "Invalid organization_id")
			}
		} else if c.Method() == "DELETE" {
			// If it is DELETE method, we need to retrieve the organization_id from the entity
			path := c.Path()
			if strings.Contains(path, "users") {
				users, err := a.userRepo.Find(&models.UserFind{ID: req.ID})
				if err != nil {
					return fiber.NewError(fiber.StatusNotFound, err.Error())
				}
				if len(*users) != 1 {
					return fiber.NewError(fiber.StatusNotFound, "User not found")
				}
				organizationID = (*users)[0].OrganizationID
			}
		}

		// Validation
		if user.OrganizationID != organizationID {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized access to organization")
		}
		return c.Next()
	}
}
