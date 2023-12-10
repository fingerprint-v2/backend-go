package services

import (
	"context"

	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
)

type UserService interface {
	GetByUsername(ctx context.Context, username string) (*models.User, error)
}

type userServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (s *userServiceImpl) GetByUsername(ctx context.Context, username string) (*models.User, error) {

	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
