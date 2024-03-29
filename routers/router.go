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
	minioHandler handlers.ObjectStorageHandler,
	organizationHandler handlers.OrganizationHandler,
	userHandler handlers.UserHandler,
	siteHandler handlers.SiteHandler,
	collectHandler handlers.LocalizeHandler,
	pointHandler handlers.PointHandler,
	mLHandler handlers.MLHandler,
	buildingHandler handlers.BuildingHandler,
	floorHandler handlers.FloorHandler,
	middleware *middleware.AuthMiddleware,
) {
	router.Get("/hello-world", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"message": "Hello World!",
		})
	})
	v1 := router.Group("/v1")
	SetUpAuthRouter(v1, validator, authHandler)
	SetUpObjectStorageRouter(v1, validator, minioHandler)
	SetupOrganizationRouter(v1, validator, organizationHandler, middleware)
	SetupSiteRouter(v1, validator, siteHandler, middleware)
	SetupUserRouter(v1, validator, userHandler, middleware)
	SetupLocalizeRouter(v1, validator, collectHandler, middleware)
	SetupPointRouter(v1, validator, pointHandler, middleware)
	SetupBuildingRouter(v1, validator, buildingHandler, middleware)
	SetupFloorRouter(v1, validator, floorHandler, middleware)
	SetupMLRouter(v1, validator, mLHandler, middleware)
}
