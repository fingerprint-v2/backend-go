package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupOrganizationRouter(router fiber.Router, v dto.Validator, handler handlers.OrganizationHandler) {
	vCreateOrganizationReq := dto.ValidateRequest[dto.CreateOrganizationReq](v)
	vUpdateOrganizationReq := dto.ValidateRequest[dto.UpdateOrganizationReq](v)
	vSearchOrganizationReq := dto.ValidateRequest[dto.SearchOrganizationReq](v)

	organization := router.Group("organizations")
	organization.Get("/:organization_id", handler.GetOrganization)
	organization.Post("/search", vSearchOrganizationReq, handler.SearchOrganization)
	organization.Put("/", vCreateOrganizationReq, handler.CreateOrganization)
	organization.Patch("/:organization_id", vUpdateOrganizationReq, handler.UpdateOrganization)
	organization.Delete("/:organization_id", handler.DeleteOrganization)
}
