package repositories

import (
	"github.com/fingerprint/models"
	"github.com/fingerprint/validates"
	"gorm.io/gorm"
)

type UserRepository interface {
	repository[models.User, validates.SearchUserReq]
	// GetByUsername(context.Context, string) (*models.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.User, validates.SearchUserReq]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.User, validates.SearchUserReq](db),
	}
}

// func (r *userRepositoryImpl) GetByUsername(ctx context.Context, username string) (*models.User, error) {
// 	user := &models.User{}
// 	if err := r.db.First(user, "username = ?", username).Error; err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }
