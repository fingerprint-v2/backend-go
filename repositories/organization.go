package repositories

import (
	"github.com/fingerprint/models"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	repository[models.Organization]
}

type organizationRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.Organization]
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepositoryImpl{
		repositoryImpl: newRepositoryImpl[models.Organization](db),
	}
}
