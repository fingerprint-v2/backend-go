package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	middleware "github.com/fingerprint/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupTrainingRouter(router fiber.Router, v dto.Validator, trainingHandler handlers.MLHandler, middleware *middleware.AuthMiddleware) {

	vCreateTrainingReq := dto.ValidateRequest[dto.CreateTrainingReq](v)

	trainingRouter := router.Group("/training")
	trainingRouter.Put("/", vCreateTrainingReq, trainingHandler.CreateTraining)

}
