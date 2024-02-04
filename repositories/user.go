package repositories

import (
	"github.com/fingerprint/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	repository[models.User, models.SearchUser]
	GetUserWithOrganization(id string) (*models.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.User, models.SearchUser]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.User, models.SearchUser](db),
	}
}

func (r *userRepositoryImpl) GetUserWithOrganization(id string) (*models.User, error) {
	user := new(models.User)
	if err := r.db.Preload("Organization").First(user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
