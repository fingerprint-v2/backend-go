package routers

import (
	"os"

	"github.com/fingerprint/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	router fiber.Router,
	organizationHandler handlers.OrganizationHandler,
	userHandler handlers.UserHandler,
) {
	router.Get("/hello-world", func(c *fiber.Ctx) error {
		return c.JSON(os.Getenv("ENV_STAGE"))
	})
	v1 := router.Group("/v1")
	SetupOrganizationRouter(v1, organizationHandler)
	SetupUserRouter(v1, userHandler)
}
