package routers

import (
	"github.com/fingerprint/handlers"
	"github.com/fingerprint/validates"
	"github.com/gofiber/fiber/v2"
)

func SetupOrganizationRouter(router fiber.Router, v validates.Validator, handler handlers.OrganizationHandler) {
	vCreateOrganizationReq := validates.ValidateRequest[validates.CreateOrganizationReq](v)
	vUpdateOrganizationReq := validates.ValidateRequest[validates.UpdateOrganizationReq](v)
	vSearchOrganizationReq := validates.ValidateRequest[validates.SearchOrganizationReq](v)

	organization := router.Group("organizations")
	organization.Get("/:organization_id", handler.GetOrganization)
	organization.Post("/search", vSearchOrganizationReq, handler.SearchOrganization)
	organization.Put("/", vCreateOrganizationReq, handler.CreateOrganization)
	organization.Patch("/:organization_id", vUpdateOrganizationReq, handler.UpdateOrganization)
	organization.Delete("/:organization_id", handler.DeleteOrganization)
}
