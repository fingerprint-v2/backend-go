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
	organization.Post("/search", middleware.SuperAdminGuard(), vSearchOrganizationReq, handler.SearchOrganization)
	organization.Put("/", middleware.SuperAdminGuard(), vCreateOrganizationReq, handler.CreateOrganization)
	organization.Patch("/", middleware.SuperAdminGuard(), vUpdateOrganizationReq, handler.UpdateOrganization)
	organization.Delete("/:organization_id", middleware.SuperAdminGuard(), handler.DeleteOrganization)
}
