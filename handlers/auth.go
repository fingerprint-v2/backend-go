package handlers

import (
	"time"

	"github.com/fingerprint/configs"
	"github.com/fingerprint/services"
	"github.com/fingerprint/utils"
	"github.com/fingerprint/validates"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
}

type authHandlerImpl struct {
	userService services.UserService
}

func NewAuthHandler(userService services.UserService) AuthHandler {
	return &authHandlerImpl{
		userService: userService,
	}
}

func checkPassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func (h *authHandlerImpl) Login(c *fiber.Ctx) error {

	req := &validates.LoginReq{}
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := h.userService.GetByUsername(c.Context(), req.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	if err := checkPassword(req.Password, user.Password); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	day := time.Hour * 24
	// Calculate JWT expiration time
	jwtExpiration := time.Now().Add(day * 1).Unix()

	// Create the JWT claims, which includes the user ID and expiry time
	claims := jwt.MapClaims{
		"user": user,
		"exp":  jwtExpiration,
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString(configs.GetAccessTokenSignature())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    t,
		MaxAge:   int(jwtExpiration - time.Now().Unix()), // Set MaxAge to match JWT expiration time
		HTTPOnly: true,
		Secure:   true,
	})

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Login success",
		Data:    nil,
	})
}
