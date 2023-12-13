package routers

import (
	"github.com/fingerprint/handlers"
	"github.com/fingerprint/validates"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRouter(router fiber.Router, handler handlers.UserHandler) {
	user := router.Group("users")
	user.Get("/me", handler.GetMe)
	user.Post("/", validates.ValidateCreateUserReq, handler.CreateUser)
	user.Put("/:user_id", validates.ValidateUpdateUserReq, handler.UpdateUser)
	user.Delete("/:user_id", handler.DeleteUser)
}
