package routers

import (
	"github.com/fingerprint/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRouter(router fiber.Router, handler handlers.UserHandler) {
	user := router.Group("users")
	// user.Get("/me", handler.GetMe)
	user.Post("/", handler.CreateUser)
	user.Put("/", handler.UpdateUser)
	user.Delete("/", handler.DeleteUser)
}
