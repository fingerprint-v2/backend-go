package handlers

import (
	"github.com/fingerprint/repositories"
	"github.com/gofiber/fiber/v2"
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

func (h *organizationHandlerImpl) GetOrganization(ctx *fiber.Ctx) error {
	return nil
}
func (h *organizationHandlerImpl) SearchOrganization(ctx *fiber.Ctx) error {
	return nil
}
func (h *organizationHandlerImpl) CreateOrganization(ctx *fiber.Ctx) error {
	return nil
}
func (h *organizationHandlerImpl) UpdateOrganization(ctx *fiber.Ctx) error {
	return nil
}
func (h *organizationHandlerImpl) DeleteOrganization(ctx *fiber.Ctx) error {
	return nil
}
