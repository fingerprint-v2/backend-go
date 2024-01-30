package routers

import (
	"github.com/fingerprint/handlers"
	"github.com/fingerprint/validates"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRouter(router fiber.Router, v validates.Validator, handler handlers.UserHandler) {
	vCreateUserReq := validates.ValidateRequest[validates.CreateUserReq](v)
	vUpdateUserReq := validates.ValidateRequest[validates.UpdateUserReq](v)

	user := router.Group("users")
	user.Get("/me", handler.GetMe)
	user.Put("/", vCreateUserReq, handler.CreateUser)
	user.Patch("/:user_id", vUpdateUserReq, handler.UpdateUser)
	user.Delete("/:user_id", handler.DeleteUser)
}
