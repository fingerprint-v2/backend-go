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

// @Tags User
// @Summary Create User
// @Description create User
// @ID create-user
// @Accept json
// @Produce json
// @Param body body validates.CreateUserReq true "Request Body"
// @Success 200 {object} utils.ResponseSuccess[string]
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/users [post]
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

// @Tags User
// @Summary Update user
// @Description update user
// @ID update-user
// @Accept json
// @Produce json
// @Param  user_id path string  true  "user's id"
// @Param body body validates.UpdateUserReq true "Request Body"
// @Success 200 {object} utils.ResponseSuccess[string]
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/users/{user_id} [put]
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

// @Tags User
// @Summary Delete User
// @Description delete User
// @ID delete-user
// @Accept json
// @Produce json
// @Param  user_id path string  true  "user's id"
// @Success 200 {object} utils.ResponseSuccess[string]
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/users/{user_id} [delete]
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
