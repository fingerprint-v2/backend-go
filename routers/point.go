package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	middleware "github.com/fingerprint/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupPointRouter(router fiber.Router, validator dto.Validator, pointHandler handlers.PointHandler, middleware *middleware.AuthMiddleware) {
	pointRouter := router.Group("/points")
	pointRouter.Post("/search", pointHandler.SearchPoint)
}
