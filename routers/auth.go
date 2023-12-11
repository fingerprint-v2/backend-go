package routers

import (
	"github.com/fingerprint/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetUpAuthRouter(router fiber.Router, handler handlers.AuthHandler) {
	router.Post("/login", handler.Login)
}
