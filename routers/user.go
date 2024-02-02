package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRouter(router fiber.Router, v dto.Validator, handler handlers.UserHandler) {
	vCreateUserReq := dto.ValidateRequest[dto.CreateUserReq](v)
	vUpdateUserReq := dto.ValidateRequest[dto.UpdateUserReq](v)

	user := router.Group("users")
	user.Get("/me", handler.GetMe)
	user.Put("/", vCreateUserReq, handler.CreateUser)
	user.Patch("/:user_id", vUpdateUserReq, handler.UpdateUser)
	user.Delete("/:user_id", handler.DeleteUser)
}
