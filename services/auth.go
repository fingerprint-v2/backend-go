package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/fingerprint/configs"
	"github.com/fingerprint/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type AuthService interface {
	ValidateToken(string) (*models.User, error)
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

	userPayload := claims["user"].(map[string]interface{})
	user := &models.User{}
	user.ID = userPayload["id"].(uuid.UUID)
	user.Username = userPayload["username"].(string)
	user.Role = userPayload["role"].(string)
	user.OrganizationID = userPayload["organization"].(string)
	return user, nil
}
