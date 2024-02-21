package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	middleware "github.com/fingerprint/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupFloorRouter(router fiber.Router, v dto.Validator, handler handlers.FloorHandler, middleware *middleware.AuthMiddleware) {

	vCreateFloorReq := dto.ValidateRequest[dto.CreateFloorReq](v)
	vSearchFloorReq := dto.ValidateRequest[dto.SearchFloorReq](v)

	floor := router.Group("floors")
	floor.Put("/", middleware.AdminGuard(), vCreateFloorReq, middleware.OrganizationGuard(), handler.CreateFloor)
	floor.Post("/search", middleware.AdminGuard(), vSearchFloorReq, middleware.OrganizationGuard(), handler.SearchFloor)
}
