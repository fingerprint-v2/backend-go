package validates

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CreateUserReq struct {
	Username       string `json:"username" validate:"required"`
	Password       string `json:"password" validate:"required"`
	Role           string `json:"role"`
	OrganizationID string `json:"organization_id" validate:"required"`
}

func ValidateCreateUserReq(c *fiber.Ctx) error {
	var req CreateUserReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := validator.New().Struct(req); err != nil {

		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Next()
}

type UpdateUserReq struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	OrganizationID string `json:"organization_id"`
}

func ValidateUpdateUserReq(c *fiber.Ctx) error {
	var req UpdateUserReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := validator.New().Struct(req); err != nil {

		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Next()
}
