package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetUpAuthRouter(router fiber.Router, v dto.Validator, handler handlers.AuthHandler) {
	vLoginReq := dto.ValidateRequest[dto.LoginReq](v)
	router.Post("/login", vLoginReq, handler.Login)
}
