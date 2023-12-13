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

// @Tags Organization
// @Summary Get Organization
// @Description get Organization
// @ID get-organization
// @Accept json
// @Produce json
// @Param  organization_id path string  true  "organization's id"
// @Success 200 {object} utils.ResponseSuccess[models.Organization]
// @Failure 400 {object} utils.ResponseError
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/organizations/{organization_id} [get]
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

// @Tags Organization
// @Summary Search Organization
// @Description search Organization
// @ID search-organization
// @Accept json
// @Produce json
// @Param body body validates.SearchOrganizationReq true "Request Body"
// @Success 200 {object} utils.ResponseSuccess[[]models.Organization]
// @Failure 400 {object} utils.ResponseError
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/organizations/search [post]
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

// @Tags Organization
// @Summary Create Organization
// @Description create Organization
// @ID create-organization
// @Accept json
// @Produce json
// @Param body body validates.CreateOrganizationReq true "Request Body"
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
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[uuid.UUID]{
		Message: "Create organization sucessfully",
		Data:    organization.ID,
	})
}

// @Tags Organization
// @Summary Update Organization
// @Description update Organization
// @ID update-organization
// @Accept json
// @Produce json
// @Param  organization_id path string  true  "organization's id"
// @Param body body validates.UpdateOrganizationReq true "Request Body"
// @Success 200 {object} utils.ResponseSuccess[string]
// @Failure 400 {object} utils.ResponseError
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/organizations/{organization_id} [put]
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
	organizationId := c.Params("organization_id")
	if err := h.organizationRepo.Delete(organizationId); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Delete organization sucessfully",
		Data:    nil,
	})
}