package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	middleware "github.com/fingerprint/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	router fiber.Router,
	validator dto.Validator,
	authHandler handlers.AuthHandler,
	minioHandler handlers.MinioHandler,
	organizationHandler handlers.OrganizationHandler,
	userHandler handlers.UserHandler,
	siteHandler handlers.SiteHandler,
	middleware *middleware.AuthMiddleware,
) {
	router.Get("/hello-world", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"message": "Hello World!",
		})
	})
	v1 := router.Group("/v1")
	SetUpAuthRouter(v1, validator, authHandler)
	SetUpMinioRouter(v1, validator, minioHandler)
	SetupOrganizationRouter(v1, validator, organizationHandler, middleware)
	SetupSiteRouter(v1, validator, siteHandler, middleware)
	SetupUserRouter(v1, validator, userHandler, middleware)
}
