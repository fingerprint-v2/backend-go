package routers

import (
	"github.com/fingerprint/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	router fiber.Router,
	authHandler handlers.AuthHandler,
	organizationHandler handlers.OrganizationHandler,
	userHandler handlers.UserHandler,
) {
	router.Get("/hello-world", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"message": "Hello World!",
		})
	})
	v1 := router.Group("/v1")
	SetupOrganizationRouter(v1, organizationHandler)
	SetupUserRouter(v1, userHandler)
	SetUpAuthRouter(v1, authHandler)
}
