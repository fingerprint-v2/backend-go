package dto

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Validator interface {
	ValidateStruct(s interface{}) error
}

type ValidatorImpl struct {
	validator *validator.Validate
}

func NewValidator() Validator {
	return &ValidatorImpl{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (v *ValidatorImpl) ValidateStruct(s interface{}) error {
	return v.validator.Struct(s)
}

func ValidateRequest[T any](v Validator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ent := new(T)
		if err := c.BodyParser(ent); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		// Check if the struct is empty
		if reflect.ValueOf(*ent).IsZero() {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", string(c.Body())))
		}

		if err := v.ValidateStruct(ent); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return c.Next()
	}
}
