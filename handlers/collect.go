package handlers

import (
	"fmt"

	"github.com/fingerprint/models"
	"github.com/fingerprint/services"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type CollectHandler interface {
	Collect(ctx *fiber.Ctx) error
}

type collectHandlerImpl struct {
	collectService services.CollectService
}

func NewCollectHandler(collectService services.CollectService) CollectHandler {
	return &collectHandlerImpl{
		collectService: collectService,
	}
}

func (h *collectHandlerImpl) Collect(c *fiber.Ctx) error {
	fmt.Println("Collect-Handler")
	err := h.collectService.Collect(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[*[]models.Organization]{
		Message: "Search collect sucessfully",
		Data:    nil,
	})
}
