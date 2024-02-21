package handlers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type SiteHandler interface {
	SearchSite(ctx *fiber.Ctx) error
	CreateSite(ctx *fiber.Ctx) error
}

type siteHandlerImpl struct {
	siteRepo         repositories.SiteRepository
	organizationRepo repositories.OrganizationRepository
}

func NewSiteHandler(siteRepo repositories.SiteRepository, organizationRepo repositories.OrganizationRepository) SiteHandler {
	return &siteHandlerImpl{
		siteRepo:         siteRepo,
		organizationRepo: organizationRepo,
	}
}

func (h *siteHandlerImpl) SearchSite(c *fiber.Ctx) error {
	req := new(dto.SearchSiteReq)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	sites, err := h.siteRepo.SearchSite(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[*[]models.Site]{
		Message: "Search site sucessfully",
		Data:    sites,
	})
}

func (h *siteHandlerImpl) CreateSite(c *fiber.Ctx) error {
	site := new(models.Site)
	if err := c.BodyParser(site); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Check if organization existed
	_, err := h.organizationRepo.Find(&models.OrganizationFind{ID: site.OrganizationID})
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Organization not found")
	}

	// Check if site already existed within same organization
	sites, err := h.siteRepo.Find(&models.SiteFind{Name: site.Name, OrganizationID: site.OrganizationID})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if len(*sites) > 0 {
		return fiber.NewError(fiber.StatusConflict, "This site name already existed within this organization")
	}

	// Creating new site
	err = h.siteRepo.Create(site)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[*models.Site]{
		Message: "Create site sucessfully",
		Data:    site,
	})
}
