package repositories

import (
	"github.com/fingerprint/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	repository[models.User]
	GetByUsername(string) (*models.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.User](db),
	}
}

func (r *userRepositoryImpl) GetByUsername(username string) (*models.User, error) {
	user := &models.User{}
	if err := r.db.First(user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return user, nil
}
