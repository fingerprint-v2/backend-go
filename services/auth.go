package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/fingerprint/configs"
	"github.com/fingerprint/constants"
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
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
	GetOrganizationIDfromContext(c *fiber.Ctx) (*string, error)
}

type authServiceImpl struct {
	userService  UserService
	userRepo     repositories.UserRepository
	siteRepo     repositories.SiteRepository
	buildingRepo repositories.BuildingRepository
	floorRepo    repositories.FloorRepository
	pointRepo    repositories.PointRepository
}

func NewAuthService(
	userService UserService,
	userRepo repositories.UserRepository,
	siteRepo repositories.SiteRepository,
	buildingRepo repositories.BuildingRepository,
	floorRepo repositories.FloorRepository,
	pointRepo repositories.PointRepository,
) AuthService {
	return &authServiceImpl{
		userService:  userService,
		userRepo:     userRepo,
		siteRepo:     siteRepo,
		buildingRepo: buildingRepo,
		floorRepo:    floorRepo,
		pointRepo:    pointRepo,
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

type organizationContextOption struct {
	OrganizationID string `json:"organization_id"`
	SiteID         string `json:"site_id"`
	BuildingID     string `json:"building_id"`
	FloorID        string `json:"floor_id"`
	ID             string `json:"id"`
}

func (s *authServiceImpl) GetOrganizationIDfromContext(c *fiber.Ctx) (*string, error) {

	req := new(organizationContextOption)
	if err := c.BodyParser(req); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	path := c.Path()
	method := c.Method()

	// Search entity. Require organization_id for all search
	if method == "POST" || strings.Contains(path, "search") {
		if req.OrganizationID == "" {
			return nil, fiber.NewError(fiber.StatusBadRequest, "No organization ID")
		}
		return &req.OrganizationID, nil
	}

	// Create and update entity
	if method == "PUT" || c.Method() == "PATCH" {
		if strings.Contains(path, "users") || strings.Contains(path, "sites") {
			if req.OrganizationID == "" {
				return nil, fiber.NewError(fiber.StatusBadRequest, "No organization ID")
			}
			return &req.OrganizationID, nil
		} else if strings.Contains(path, "buildings") {
			if req.SiteID == "" {
				return nil, fiber.NewError(fiber.StatusBadRequest, "No site ID")
			}
			site, err := s.siteRepo.Get(req.SiteID)
			if err != nil {
				return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return &site.OrganizationID, nil

		} else if strings.Contains(path, "floors") {
			if req.BuildingID == "" {
				return nil, fiber.NewError(fiber.StatusBadRequest, "No building ID")
			}
			building, err := s.buildingRepo.Get(req.BuildingID)
			if err != nil {
				return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return &building.OrganizationID, nil
		} else if strings.Contains(path, "points") {
			if req.FloorID == "" {
				return nil, fiber.NewError(fiber.StatusBadRequest, "No floor ID")
			}
			floor, err := s.floorRepo.Get(req.FloorID)
			if err != nil {
				return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return &floor.OrganizationID, nil
		}
	}

	// Delete entity
	if method == "DELETE" {

		if req.ID == "" {
			return nil, fiber.NewError(fiber.StatusBadRequest, "No ID")
		}

		if strings.Contains(path, "users") {
			user, err := s.userRepo.Get(req.ID)
			if err != nil {
				return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return &user.OrganizationID, nil
		} else if strings.Contains(path, "sites") {
			site, err := s.siteRepo.Get(req.ID)
			if err != nil {
				return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return &site.OrganizationID, nil
		} else if strings.Contains(path, "buildings") {
			building, err := s.buildingRepo.Get(req.ID)
			if err != nil {
				return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return &building.OrganizationID, nil
		} else if strings.Contains(path, "floors") {
			floor, err := s.floorRepo.Get(req.ID)
			if err != nil {
				return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return &floor.OrganizationID, nil
		} else if strings.Contains(path, "points") {
			point, err := s.pointRepo.Get(req.ID)
			if err != nil {
				return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
			}
			return &point.OrganizationID, nil
		}
	}

	return nil, fiber.NewError(fiber.StatusBadRequest, "Cannot obtain organization context")
}
