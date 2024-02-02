package handlers

import (
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/services"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler interface {
	GetMe(*fiber.Ctx) error
	CreateUser(*fiber.Ctx) error
	UpdateUser(*fiber.Ctx) error
	DeleteUser(*fiber.Ctx) error
}

type userHandlerImpl struct {
	authService services.AuthService
	userRepo    repositories.UserRepository
}

func NewUserHandler(authService services.AuthService, userRepo repositories.UserRepository) UserHandler {
	return &userHandlerImpl{
		authService: authService,
		userRepo:    userRepo,
	}
}

// @Tags User
// @Summary Get Me
// @Description get Me
// @ID get-me
// @Accept json
// @Produce json
// @Param body body dto.CreateUserReq true "Request Body"
// @Success 200 {object} utils.ResponseSuccess[models.User]
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/users/me [post]
func (h *userHandlerImpl) GetMe(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.User)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid User")
	}

	user, err := h.userRepo.Get(c.Context(), user.ID.String())

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[models.User]{
		Message: "Get me sucessfully",
		Data:    *user,
	})
}

// @Tags User
// @Summary Create User
// @Description create User
// @ID create-user
// @Accept json
// @Produce json
// @Param body body dto.CreateUserReq true "Request Body"
// @Success 200 {object} utils.ResponseSuccess[string]
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/users [post]
func (h *userHandlerImpl) CreateUser(c *fiber.Ctx) error {
	ctx := c.Context()
	user := &models.User{
		ID: uuid.New(),
	}
	if err := c.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := h.authService.HashPassword(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := h.userRepo.Create(ctx, user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[*models.User]{
		Message: "Create user sucessfully",
		Data:    user,
	})
}

// @Tags User
// @Summary Update user
// @Description update user
// @ID update-user
// @Accept json
// @Produce json
// @Param  user_id path string  true  "user's id"
// @Param body body dto.UpdateUserReq true "Request Body"
// @Success 200 {object} utils.ResponseSuccess[string]
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/users/{user_id} [put]
func (h *userHandlerImpl) UpdateUser(c *fiber.Ctx) error {
	ctx := c.Context()
	userId := c.Params("user_id")
	user := &models.User{}
	if err := c.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := h.authService.HashPassword(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := h.userRepo.Update(ctx, userId, user); err != nil {
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
	ctx := c.Context()
	userId := c.Params("user_id")
	if err := h.userRepo.Delete(ctx, userId); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Delete user sucessfully",
		Data:    nil,
	})
}
