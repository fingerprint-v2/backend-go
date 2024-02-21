package handlers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type FloorHandler interface {
	CreateFloor(c *fiber.Ctx) error
	SearchFloor(c *fiber.Ctx) error
}

type floorHandlerImpl struct {
	floorRepo repositories.FloorRepository
	building  repositories.BuildingRepository
}

func NewFloorHandler(
	floorRepo repositories.FloorRepository,
	building repositories.BuildingRepository,
) FloorHandler {
	return &floorHandlerImpl{
		floorRepo: floorRepo,
		building:  building,
	}
}

func (h *floorHandlerImpl) SearchFloor(c *fiber.Ctx) error {
	req := new(dto.SearchFloorReq)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	floors, err := h.floorRepo.SearchFloor(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(
		utils.ResponseSuccess[*[]models.Floor]{
			Message: "Search floor successfully",
			Data:    floors,
		})

}

func (h *floorHandlerImpl) CreateFloor(c *fiber.Ctx) error {

	floor := new(models.Floor)
	if err := c.BodyParser(floor); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Check if building exists
	building, err := h.building.Get(floor.BuildingID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Building not found")
	}

	floor.BuildingID = building.ID.String()
	floor.SiteID = building.SiteID
	floor.OrganizationID = building.OrganizationID

	// Check if floor exists in the same building
	floors, err := h.floorRepo.Find(&models.FloorFind{Name: floor.Name, BuildingID: floor.BuildingID})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if len(*floors) > 0 {
		return fiber.NewError(fiber.StatusConflict, "Building already exists within this site")
	}

	// Create floor
	err = h.floorRepo.Create(floor)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(
		utils.ResponseSuccess[*models.Floor]{
			Message: "Floor created successfully",
			Data:    floor,
		})
}
