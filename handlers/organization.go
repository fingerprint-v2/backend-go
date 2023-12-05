package handlers

import (
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
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
	organizationRepo repositories.OrganizationRepository
}

func NewOrganizationHandler(organizationRepo repositories.OrganizationRepository) OrganizationHandler {
	return &organizationHandlerImpl{
		organizationRepo: organizationRepo,
	}
}

func (h *organizationHandlerImpl) GetOrganization(c *fiber.Ctx) error {
	organizationId := c.Params("organization_id")
	organization, err := h.organizationRepo.Get(organizationId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[*models.Organization]{
		Message: "Get organization sucessfully",
		Data:    organization,
	})
}

func (h *organizationHandlerImpl) SearchOrganization(c *fiber.Ctx) error {
	return nil
}

func (h *organizationHandlerImpl) CreateOrganization(c *fiber.Ctx) error {
	organization := &models.Organization{}
	if err := c.BodyParser(organization); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	ent, err := h.organizationRepo.Create(organization)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[uuid.UUID]{
		Message: "Create organization sucessfully",
		Data:    ent.ID,
	})
}

func (h *organizationHandlerImpl) UpdateOrganization(c *fiber.Ctx) error {
	organizationId := c.Params("organization_id")
	organization := &models.Organization{}
	if err := c.BodyParser(organization); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	ent, err := h.organizationRepo.Update(organizationId, organization)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[uuid.UUID]{
		Message: "Update organization sucessfully",
		Data:    ent.ID,
	})
}

func (h *organizationHandlerImpl) DeleteOrganization(c *fiber.Ctx) error {
	organizationId := c.Params("organization_id")
	if err := h.organizationRepo.Delete(organizationId); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Update organization sucessfully",
		Data:    nil,
	})
}
