package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	middleware "github.com/fingerprint/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupLocalizeRouter(router fiber.Router, v dto.Validator, handler handlers.LocalizeHandler, middleware *middleware.AuthMiddleware) {

	vCreateSupervisedSurveyReq := dto.ValidateRequest[dto.CreateSupervisedSurveyReq](v)
	vCreateUnsupervisedSurveyReq := dto.ValidateRequest[dto.CreateUnsupervisedSurveyReq](v)

	localize := router.Group("localize")
	localize.Put("/supervised", vCreateSupervisedSurveyReq, handler.CreateSupervisedSurvey)
	localize.Put("/unsupervised", vCreateUnsupervisedSurveyReq, handler.CreateUnsupervisedSurvey)
}
