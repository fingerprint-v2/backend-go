package handlers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/services"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type CollectHandler interface {
	CreateSurvey(ctx *fiber.Ctx) error
}

type collectHandlerImpl struct {
	collectService services.CollectService
}

func NewCollectHandler(collectService services.CollectService) CollectHandler {
	return &collectHandlerImpl{
		collectService: collectService,
	}
}

func (h *collectHandlerImpl) CreateSurvey(c *fiber.Ctx) error {
	req := new(dto.CreateSurveyReq)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := h.collectService.Collect(c.Context(), req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[*[]models.Organization]{
		Message: "Search collect sucessfully",
		Data:    nil,
	})
}
