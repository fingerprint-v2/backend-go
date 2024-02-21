package handlers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/services"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type MLHandler interface {
	CreateTraining(c *fiber.Ctx) error
}

type mLHandlerImpl struct {
	trainingService services.MLService
}

func NewMLHandler(trainingService services.MLService) MLHandler {
	return &mLHandlerImpl{
		trainingService: trainingService,
	}
}

func (h *mLHandlerImpl) CreateTraining(c *fiber.Ctx) error {

	req := new(dto.CreateTrainingReq)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := h.trainingService.CreateTraining(c, req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Create training successfully",
		Data:    "Training is being processed, please wait for a moment.",
	})

}
