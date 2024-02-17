package handlers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/services"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type TrainingHandler interface {
	CreateTraining(c *fiber.Ctx) error
}

type trainingHandlerImpl struct {
	trainingService services.TrainingService
}

func NewTrainingHandler(trainingService services.TrainingService) TrainingHandler {
	return &trainingHandlerImpl{
		trainingService: trainingService,
	}
}

func (h *trainingHandlerImpl) CreateTraining(c *fiber.Ctx) error {

	req := new(dto.CreateTrainingReq)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.trainingService.CreateTraining(c, req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Create training successfully",
		Data:    nil,
	})

}
