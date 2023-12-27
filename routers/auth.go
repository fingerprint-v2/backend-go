package routers

import (
	"github.com/fingerprint/handlers"
	"github.com/fingerprint/validates"
	"github.com/gofiber/fiber/v2"
)

func SetUpAuthRouter(router fiber.Router, v *validates.Validator, handler handlers.AuthHandler) {
	vLoginReq := validates.ValidateRequest[validates.LoginReq](v)
	router.Post("/login", vLoginReq, handler.Login)
}
