package handlers

import (
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/services"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrganizationHandler interface {
	GetOrganization(ctx *fiber.Ctx) error
	SearchOrganization(ctx *fiber.Ctx) error
	CreateOrganization(ctx *fiber.Ctx) error
	UpdateOrganization(ctx *fiber.Ctx) error
	DeleteOrganization(ctx *fiber.Ctx) error
}

type organizationHandlerImpl struct {
	organizationRepo    repositories.OrganizationRepository
	organizationService services.OrganizationService
}

func NewOrganizationHandler(organizationService services.OrganizationService, organizationRepo repositories.OrganizationRepository) OrganizationHandler {
	return &organizationHandlerImpl{
		organizationService: organizationService,
		organizationRepo:    organizationRepo,
	}
}

func (h *organizationHandlerImpl) GetOrganization(c *fiber.Ctx) error {
	organizationId := c.Params("organization_id")
	organization, err := h.organizationRepo.Get(organizationId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[models.Organization]{
		Message: "Get organization sucessfully",
		Data:    *organization,
	})
}

func (h *organizationHandlerImpl) SearchOrganization(c *fiber.Ctx) error {
	ctx := c.Context()
	organization := &models.Organization{}
	if err := c.BodyParser(organization); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	organizations, err := h.organizationService.SearchOrganization(ctx, organization)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[[]models.Organization]{
		Message: "Search organization sucessfully",
		Data:    organizations,
	})
}

func (h *organizationHandlerImpl) CreateOrganization(c *fiber.Ctx) error {
	organization := &models.Organization{
		ID: uuid.New(),
	}
	if err := c.BodyParser(organization); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := h.organizationRepo.Create(organization); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[uuid.UUID]{
		Message: "Create organization sucessfully",
		Data:    organization.ID,
	})
}

func (h *organizationHandlerImpl) UpdateOrganization(c *fiber.Ctx) error {
	organizationId := c.Params("organization_id")
	organization := &models.Organization{}
	if err := c.BodyParser(organization); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := h.organizationRepo.Update(organizationId, organization); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Update organization sucessfully",
		Data:    nil,
	})
}

func (h *organizationHandlerImpl) DeleteOrganization(c *fiber.Ctx) error {
	organizationId := c.Params("organization_id")
	if err := h.organizationRepo.Delete(organizationId); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Delete organization sucessfully",
		Data:    nil,
	})
}
