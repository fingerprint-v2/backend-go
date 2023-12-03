package routers

import (
	"github.com/fingerprint/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupOrganizationRouter(router fiber.Router, handler handlers.OrganizationHandler) {
	organization := router.Group("organizations")
	organization.Get("/", handler.GetOrganization)
	organization.Post("/search", handler.SearchOrganization)
	organization.Post("/", handler.CreateOrganization)
	organization.Put("/", handler.UpdateOrganization)
	organization.Delete("/", handler.DeleteOrganization)
}
