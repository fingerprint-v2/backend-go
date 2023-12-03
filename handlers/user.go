package handlers

import (
	"github.com/fingerprint/repositories"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	CreateUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
}

type userHandlerImpl struct {
	userRepo repositories.UserRepository
}

func NewUserHandler(userRepo repositories.UserRepository) UserHandler {
	return &userHandlerImpl{
		userRepo: userRepo,
	}
}

func (h *userHandlerImpl) CreateUser(ctx *fiber.Ctx) error {
	return nil
}
func (h *userHandlerImpl) UpdateUser(ctx *fiber.Ctx) error {
	return nil
}
func (h *userHandlerImpl) DeleteUser(ctx *fiber.Ctx) error {
	return nil
}
