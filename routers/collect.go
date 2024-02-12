package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	middleware "github.com/fingerprint/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupCollectRouter(router fiber.Router, v dto.Validator, handler handlers.CollectHandler, middleware *middleware.AuthMiddleware) {

	vCreateSurveyReq := dto.ValidateRequest[dto.CreateSurveyReq](v)

	collect := router.Group("collect")
	collect.Post("/", vCreateSurveyReq, handler.CreateSurvey)
}
