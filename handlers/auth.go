package handlers

import (
	"time"

	"github.com/fingerprint/constants"
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/services"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
}

type authHandlerImpl struct {
	authService services.AuthService
	userRepo    repositories.UserRepository
}

func NewAuthHandler(authService services.AuthService, userRepo repositories.UserRepository) AuthHandler {
	return &authHandlerImpl{
		authService: authService,
		userRepo:    userRepo,
	}
}

// @Tags Auth
// @Summary Login
// @Description login
// @ID login
// @Accept json
// @Produce json
// @Param body body dto.LoginReq true "Request Body"
// @Success 200 {object} utils.ResponseSuccess[string]
// @Failure 400 {object} utils.ResponseError
// @Failure 401 {object} utils.ResponseError
// @Failure 404 {object} utils.ResponseError
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/login [post]
func (h *authHandlerImpl) Login(c *fiber.Ctx) error {

	req := &dto.LoginReq{}
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	users, err := h.userRepo.Search(c.Context(), &models.SearchUser{Username: req.Username})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if len(*users) != 1 {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}
	// Extract user pointer
	user := &(*users)[0]

	if err := h.authService.CheckPassword(req.Password, user.Password); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	t, err := h.authService.GenerateToken(user)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *t,
		MaxAge:   int(int64(constants.JWTExpiration) - time.Now().Unix()),
		HTTPOnly: true,
		Secure:   true,
	})

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Login successful",
		Data:    nil,
	})
}
