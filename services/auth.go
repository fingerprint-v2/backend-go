package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/fingerprint/configs"
	"github.com/fingerprint/constants"
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	ValidateToken(string) (*models.User, error)
	GenerateToken(*models.User) (*string, error)
	CheckPassword(string, string) error
	HashPassword(user *models.User) error
	CheckValidRole(string) error
}

type authServiceImpl struct {
	userService UserService
}

func NewAuthService(userService UserService) AuthService {
	return &authServiceImpl{
		userService: userService,
	}
}

func parseToken(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return configs.GetAccessTokenSignature(), nil
	})
	if err != nil {
		return nil, err
	}
	return t, nil
}

func isTokenExpired(token string) (bool, error) {
	t, err := parseToken(token)
	if err != nil {
		return false, err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		expirationTime := int64(claims["exp"].(float64))
		currentTime := time.Now().Unix()
		return currentTime > expirationTime, nil
	}

	return false, errors.New("invalid Token")
}

func (s *authServiceImpl) ValidateToken(token string) (*models.User, error) {
	t, err := parseToken(token)
	if err != nil {
		return nil, err
	}

	isExpired, err := isTokenExpired(token)
	if err != nil {
		return nil, err
	}
	if isExpired {
		return nil, errors.New("the token has expired")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot extract jwt payload")
	}

	userPayload, ok := claims["user"].(map[string]interface{})
	if !ok {
		return nil, errors.New("user data is not in the expected format")
	}

	return utils.ConvertPayloadToUser(userPayload)
}

func (s *authServiceImpl) CheckPassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func (s *authServiceImpl) GenerateToken(user *models.User) (*string, error) {

	// Convert user to userCookie
	userCookie, err := utils.TypeConverter[dto.CookiePayloadUser](user)

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	claims := jwt.MapClaims{
		"user": userCookie,
		"exp":  constants.JWTExpiration,
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(configs.GetAccessTokenSignature())
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &t, nil
}

func (s *authServiceImpl) HashPassword(user *models.User) error {
	if user.Password == "" {
		return errors.New("password is required")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (s *authServiceImpl) CheckValidRole(role string) error {

	roles := []constants.UserRole{constants.SUPERADMIN, constants.ADMIN, constants.USER}
	for _, value := range roles {
		if role == value.String() {
			return nil
		}
	}
	return errors.New("invalid role")
}
