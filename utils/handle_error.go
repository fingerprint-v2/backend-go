package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func HandleError(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	messageError := fiber.ErrInternalServerError.Message
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	switch e.Code {
	case 400:
		messageError = fiber.ErrBadRequest.Message
	case 401:
		messageError = fiber.ErrUnauthorized.Message
	case 404:
		messageError = fiber.ErrNotFound.Message
	case 409:
		messageError = fiber.ErrConflict.Message
	}

	return c.Status(code).JSON(fiber.Map{
		"error": messageError,
	})
}
