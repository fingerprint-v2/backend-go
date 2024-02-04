package repositories

import (
	"errors"

	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	repository[models.User, models.UserFind]
	SearchUser(*dto.SearchUserReq) (*[]models.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.User, models.UserFind]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.User, models.UserFind](db),
	}
}

func (r *userRepositoryImpl) SearchUser(req *dto.SearchUserReq) (*[]models.User, error) {

	users := new([]models.User)

	// Construct the map from the request
	userMap, err := utils.TypeConverter[map[string]interface{}](req)
	if err != nil {
		return nil, err
	}
	delete(*userMap, "with_organization")
	delete(*userMap, "all")
	// Make sure that map is not empty when "all" is false
	if len(*userMap) == 0 && !req.All {
		return nil, errors.New("no valid search parameters provided")
	}
	if req.All {
		userMap = &map[string]interface{}{}
	}

	// Optional preload
	db := r.db
	if req.WithOrganization {
		db = db.Preload("Organization")
	}

	// DB query
	if err := db.Find(users, *userMap).Error; err != nil {
		return nil, err
	}
	return users, nil
}
