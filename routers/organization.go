package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	middleware "github.com/fingerprint/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupOrganizationRouter(router fiber.Router, v dto.Validator, handler handlers.OrganizationHandler, middleware *middleware.AuthMiddleware) {
	vCreateOrganizationReq := dto.ValidateRequest[dto.CreateOrganizationReq](v)
	vUpdateOrganizationReq := dto.ValidateRequest[dto.UpdateOrganizationReq](v)
	vSearchOrganizationReq := dto.ValidateRequest[dto.SearchOrganizationReq](v)

	organization := router.Group("organizations")
	organization.Get("/:organization_id", handler.GetOrganization)
	organization.Post("/search", vSearchOrganizationReq, handler.SearchOrganization)
	organization.Put("/", middleware.AdminGuard(), vCreateOrganizationReq, handler.CreateOrganization)
	organization.Patch("/:organization_id", middleware.AdminGuard(), vUpdateOrganizationReq, handler.UpdateOrganization)
	organization.Delete("/:organization_id", middleware.AdminGuard(), handler.DeleteOrganization)
}
