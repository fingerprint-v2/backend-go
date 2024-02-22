package handlers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/services"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type LocalizeHandler interface {
	CreateSupervisedSurvey(ctx *fiber.Ctx) error
	CreateUnsupervisedSurvey(ctx *fiber.Ctx) error
}

type localizeHandlerImpl struct {
	localizeService services.LocalizeService
}

func NewLocalizeHandler(localizeService services.LocalizeService) LocalizeHandler {
	return &localizeHandlerImpl{
		localizeService: localizeService,
	}
}

func (h *localizeHandlerImpl) CreateSupervisedSurvey(c *fiber.Ctx) error {
	req := new(dto.CreateSurpervisedSurveyReq)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user := c.Locals("user").(*models.User)
	err := h.localizeService.CreateSupervisedSurvey(req, user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Create survey successfully",
		Data:    nil,
	})
}

func (h *localizeHandlerImpl) CreateUnsupervisedSurvey(c *fiber.Ctx) error {
	req := new(dto.CreateUnsurpervisedSurveyReq)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user := c.Locals("user").(*models.User)
	err := h.localizeService.CreateUnsupervisedSurvey(req, user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Create survey successfully",
		Data:    nil,
	})
}
