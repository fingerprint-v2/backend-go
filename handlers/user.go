package handlers

import (
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func hashPassword(user *models.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (h *userHandlerImpl) CreateUser(c *fiber.Ctx) error {
	user := &models.User{
		ID: uuid.New(),
	}
	if err := c.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := hashPassword(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := h.userRepo.Create(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[uuid.UUID]{
		Message: "Create user sucessfully",
		Data:    user.ID,
	})
}

func (h *userHandlerImpl) UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	user := &models.User{}
	if err := c.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := hashPassword(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := h.userRepo.Update(userId, user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Update user sucessfully",
		Data:    nil,
	})
}
func (h *userHandlerImpl) DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	if err := h.userRepo.Delete(userId); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Delete user sucessfully",
		Data:    nil,
	})
}
