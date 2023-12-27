package validates

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func ValidateRequest[T any](v *Validator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ent := new(T)
		if err := c.BodyParser(ent); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if err := v.validator.Struct(ent); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.Next()
	}
}
