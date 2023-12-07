package validates

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateOrganizationReq struct {
	Name string `json:"name" validate:"required"`
}

func ValidateCreateOrganizationReq(c *fiber.Ctx) error {
	var req CreateOrganizationReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := validator.New().Struct(req); err != nil {

		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Next()
}

type SearchOrganizationReq struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func ValidateSearchOrganizationReq(c *fiber.Ctx) error {
	var req SearchOrganizationReq
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := validator.New().Struct(req); err != nil {

		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Next()
}
