package handlers

import "github.com/gofiber/fiber/v2"

type TrainingHandler interface {
	CreateTraining(ctx *fiber.Ctx) error
}

type trainingHandlerImpl struct{}

func NewTrainingHandler() TrainingHandler {
	return &trainingHandlerImpl{}
}

func (h *trainingHandlerImpl) CreateTraining(ctx *fiber.Ctx) error {
	return nil
}
