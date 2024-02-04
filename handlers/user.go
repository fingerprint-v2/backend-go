package handlers

import (
	"github.com/fingerprint/dto"
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
	SearchUser(*fiber.Ctx) error
}

type userHandlerImpl struct {
	authService      services.AuthService
	userRepo         repositories.UserRepository
	organizationRepo repositories.OrganizationRepository
}

func NewUserHandler(
	authService services.AuthService,
	userRepo repositories.UserRepository,
	organizationRepo repositories.OrganizationRepository,
) UserHandler {
	return &userHandlerImpl{
		authService:      authService,
		userRepo:         userRepo,
		organizationRepo: organizationRepo,
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

	users, err := h.userRepo.SearchUser(&dto.SearchUserReq{ID: user.ID.String(), WithOrganization: true})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if len(*users) != 1 {
		return fiber.NewError(fiber.StatusInternalServerError, "Invalid User")
	}
	user = &(*users)[0]

	// Remove password from response
	user.Password = ""

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
	user := &models.User{
		ID: uuid.New(),
	}
	if err := c.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Check for valid role
	if err := h.authService.CheckValidRole(user.Role); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Check for valid organization (not really needed because of data integrity)
	orgs, err := h.organizationRepo.Find(&models.OrganizationFind{ID: user.OrganizationID})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if len(*orgs) != 1 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid organization ID")
	}

	if err := h.authService.HashPassword(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := h.userRepo.Create(user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Remove password from response
	user.Password = ""

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

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if user.Password != "" {
		if err := h.authService.HashPassword(user); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	if err := h.authService.CheckValidRole(user.Role); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := h.userRepo.Update(user.ID.String(), user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Remove password from response
	user.Password = ""

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[*models.User]{
		Message: "Update user sucessfully",
		Data:    user,
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

func (h *userHandlerImpl) SearchUser(c *fiber.Ctx) error {

	user := new(dto.SearchUserReq)
	if err := c.BodyParser(user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	users, err := h.userRepo.SearchUser(user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Remove password from response
	for i := range *users {
		(*users)[i].Password = ""
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[*[]models.User]{
		Message: "Search user sucessfully",
		Data:    users,
	})
}
