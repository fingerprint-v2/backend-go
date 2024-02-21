package handlers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type PointHandler interface {
	CreatePoint(ctx *fiber.Ctx) error
	SearchPoint(ctx *fiber.Ctx) error
}

type pointHandlerImpl struct {
	pointRepo repositories.PointRepository
	floorRepo repositories.FloorRepository
}

func NewPointHandler(pointRepo repositories.PointRepository, floorRepo repositories.FloorRepository) PointHandler {
	return &pointHandlerImpl{
		pointRepo: pointRepo,
		floorRepo: floorRepo,
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

func (h *pointHandlerImpl) CreatePoint(c *fiber.Ctx) error {
	point := new(models.Point)
	if err := c.BodyParser(point); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Check if floor exists
	floor, err := h.floorRepo.Get(point.FloorID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Floor not found")
	}

	point.FloorID = floor.ID.String()
	point.BuildingID = floor.BuildingID
	point.SiteID = floor.SiteID
	point.OrganizationID = floor.OrganizationID

	// Check if point exists
	points, err := h.pointRepo.Find(&models.PointFind{
		Name:   point.Name,
		SiteID: point.SiteID,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if len(*points) > 0 {
		return fiber.NewError(fiber.StatusConflict, "Building already exists within this site")
	}

	// Create point
	err = h.pointRepo.Create(point)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(
		utils.ResponseSuccess[*models.Point]{
			Message: "Create point successfully",
			Data:    point,
		})

}
