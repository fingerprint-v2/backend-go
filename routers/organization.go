package routers

import (
	"github.com/fingerprint/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupOrganizationRouter(router fiber.Router, handler handlers.OrganizationHandler) {
	organization := router.Group("organizations")
	organization.Get("/:organization_id", handler.GetOrganization)
	organization.Post("/search", handler.SearchOrganization)
	organization.Post("/", handler.CreateOrganization)
	organization.Put("/:organization_id", handler.UpdateOrganization)
	organization.Delete("/:organization_id", handler.DeleteOrganization)
}
