package handlers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type PointHandler interface {
	SearchPoint(ctx *fiber.Ctx) error
}

type pointHandlerImpl struct {
	pointRepo repositories.PointRepository
}

func NewPointHandler(pointRepo repositories.PointRepository) PointHandler {
	return &pointHandlerImpl{
		pointRepo: pointRepo,
	}
}

func (h *pointHandlerImpl) SearchPoint(c *fiber.Ctx) error {

	req := new(dto.SearchPointReq)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	points, err := h.pointRepo.SearchPoint(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[*[]models.Point]{
		Message: "Search point sucessfully",
		Data:    points,
	})

}
