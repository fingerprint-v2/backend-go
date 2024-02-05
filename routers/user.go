package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	middleware "github.com/fingerprint/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRouter(router fiber.Router, v dto.Validator, handler handlers.UserHandler, middleware *middleware.AuthMiddleware) {
	vCreateUserReq := dto.ValidateRequest[dto.CreateUserReq](v)
	vUpdateUserReq := dto.ValidateRequest[dto.UpdateUserReq](v)
	vSearchUserReq := dto.ValidateRequest[dto.SearchUserReq](v)
	vDeleteUserReq := dto.ValidateRequest[dto.DeleteUserReq](v)

	user := router.Group("users")
	user.Get("/me", handler.GetMe)
	user.Put("/", middleware.AdminGuard(), vCreateUserReq, handler.CreateUser)
	user.Patch("/", middleware.AdminGuard(), vUpdateUserReq, handler.UpdateUser)
	user.Delete("/", middleware.AdminGuard(), vDeleteUserReq, handler.DeleteUser)
	user.Post("/search", middleware.AdminGuard(), vSearchUserReq, handler.SearchUser)
}
