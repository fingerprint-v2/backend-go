package handlers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrganizationHandler interface {
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

// @Tags Organization
// @Summary Search Organization
// @Description search Organization
// @ID search-organization
// @Accept json
// @Produce json
// @Param body body dto.SearchOrganizationReq true "Request Body"
// @Success 200 {object} utils.ResponseSuccess[[]models.Organization]
// @Failure 400 {object} utils.ResponseError
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/organizations/search [post]
func (h *organizationHandlerImpl) SearchOrganization(c *fiber.Ctx) error {
	req := new(dto.SearchOrganizationReq)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	organizations, err := h.organizationRepo.SearchOrganization(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[*[]models.Organization]{
		Message: "Search organization sucessfully",
		Data:    organizations,
	})
}

// @Tags Organization
// @Summary Create Organization
// @Description create Organization
// @ID create-organization
// @Accept json
// @Produce json
// @Param body body dto.CreateOrganizationReq true "Request Body"
// @Success 200 {object} utils.ResponseSuccess[uuid.UUID]
// @Failure 400 {object} utils.ResponseError
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/organizations [post]
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
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[*models.Organization]{
		Message: "Create organization sucessfully",
		Data:    organization,
	})
}

// @Tags Organization
// @Summary Update Organization
// @Description update Organization
// @ID update-organization
// @Accept json
// @Produce json
// @Param  organization_id path string  true  "organization's id"
// @Param body body dto.UpdateOrganizationReq true "Request Body"
// @Success 200 {object} utils.ResponseSuccess[string]
// @Failure 400 {object} utils.ResponseError
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/organizations/{organization_id} [put]
func (h *organizationHandlerImpl) UpdateOrganization(c *fiber.Ctx) error {

	org := new(models.Organization)
	if err := c.BodyParser(org); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := h.organizationRepo.Update(org.ID.String(), org); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[*models.Organization]{
		Message: "Update organization sucessfully",
		Data:    org,
	})
}

// @Tags Organization
// @Summary Delete Organization
// @Description delete Organization
// @ID delete-organization
// @Accept json
// @Produce json
// @Param  organization_id path string  true  "organization's id"
// @Success 200 {object} utils.ResponseSuccess[string]
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/organizations/{organization_id} [delete]
func (h *organizationHandlerImpl) DeleteOrganization(c *fiber.Ctx) error {
	req := new(dto.DeleteOrganizationReq)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.organizationRepo.Delete(req.ID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Delete organization sucessfully",
		Data:    nil,
	})
}
