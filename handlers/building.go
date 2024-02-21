package handlers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type BuildingHandler interface {
	SearchBuilding(ctx *fiber.Ctx) error
	CreateBuilding(ctx *fiber.Ctx) error
}

type buildingHandlerImpl struct {
	buildingRepo repositories.BuildingRepository
	siteRepo     repositories.SiteRepository
}

func NewBuildingHandler(
	buildingRepo repositories.BuildingRepository,
	siteRepo repositories.SiteRepository,
) BuildingHandler {
	return &buildingHandlerImpl{
		buildingRepo: buildingRepo,
		siteRepo:     siteRepo,
	}
}

func (h *buildingHandlerImpl) SearchBuilding(c *fiber.Ctx) error {

	req := new(dto.SearchBuildingReq)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	buildings, err := h.buildingRepo.SearchBuilding(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[*[]models.Building]{
		Message: "Search buildings successfully",
		Data:    buildings,
	})
}

func (h *buildingHandlerImpl) CreateBuilding(ctx *fiber.Ctx) error {
	building := new(models.Building)
	if err := ctx.BodyParser(building); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Check if the site exists
	site, err := h.siteRepo.Get(building.SiteID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Site not found")
	}

	// Add organization ID to the building
	building.OrganizationID = site.OrganizationID

	// Check if the building already exists in the same site

	buildings, err := h.buildingRepo.Find(&models.BuildingFind{Name: building.Name, SiteID: building.SiteID})

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if len(*buildings) > 0 {
		return fiber.NewError(fiber.StatusConflict, "Building already exists within this site")
	}

	// Creating new building
	err = h.buildingRepo.Create(building)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(utils.ResponseSuccess[*models.Building]{
		Message: "Building created successfully",
		Data:    building,
	})
}
