package routers

import (
	"github.com/fingerprint/handlers"
	"github.com/fingerprint/validates"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	router fiber.Router,
	validator *validates.Validator,
	authHandler handlers.AuthHandler,
	minioHandler handlers.MinioHandler,
	organizationHandler handlers.OrganizationHandler,
	userHandler handlers.UserHandler,
) {
	router.Get("/hello-world", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"message": "Hello World!",
		})
	})
	v1 := router.Group("/v1")
	SetUpAuthRouter(v1, validator, authHandler)
	SetUpMinioRouter(v1, validator, minioHandler)
	SetupOrganizationRouter(v1, validator, organizationHandler)
	SetupUserRouter(v1, validator, userHandler)
}
