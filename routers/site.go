package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	middleware "github.com/fingerprint/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupSiteRouter(router fiber.Router, v dto.Validator, handler handlers.SiteHandler, middleware *middleware.AuthMiddleware) {
	vCreateSiteReq := dto.ValidateRequest[dto.CreateSiteReq](v)
	vSearchSiteReq := dto.ValidateRequest[dto.SearchSiteReq](v)

	site := router.Group("sites")
	site.Put("/", middleware.AdminGuard(), vCreateSiteReq, middleware.OrganizationGuard(), handler.CreateSite)
	site.Post("/search", middleware.AdminGuard(), vSearchSiteReq, middleware.OrganizationGuard(), handler.SearchSite)
}
