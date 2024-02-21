package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	middleware "github.com/fingerprint/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupPointRouter(router fiber.Router, v dto.Validator, handler handlers.PointHandler, middleware *middleware.AuthMiddleware) {

	vCreatePointReq := dto.ValidateRequest[dto.CreatePointReq](v)
	vSearchPointReq := dto.ValidateRequest[dto.SearchPointReq](v)

	point := router.Group("points")
	point.Put("/", middleware.AdminGuard(), vCreatePointReq, middleware.OrganizationGuard(), handler.CreatePoint)
	point.Post("/search", middleware.AdminGuard(), vSearchPointReq, middleware.OrganizationGuard(), handler.SearchPoint)
}
