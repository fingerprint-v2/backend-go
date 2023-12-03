package repositories

import (
	"github.com/fingerprint/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	repository[models.User]
}

type userRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{}
}
