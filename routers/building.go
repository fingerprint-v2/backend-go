package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	middleware "github.com/fingerprint/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupBuildingRouter(router fiber.Router, v dto.Validator, handler handlers.BuildingHandler, middleware *middleware.AuthMiddleware) {

	vCreateBuildingReq := dto.ValidateRequest[dto.CreateBuildingReq](v)
	vSearchBuildingReq := dto.ValidateRequest[dto.SearchBuildingReq](v)

	building := router.Group("buildings")
	building.Put("/", middleware.AdminGuard(), vCreateBuildingReq, middleware.OrganizationGuard(), handler.CreateBuilding)
	building.Post("/search", middleware.AdminGuard(), vSearchBuildingReq, middleware.OrganizationGuard(), handler.SearchBuilding)
}
