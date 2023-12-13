package repositories

import (
	"github.com/fingerprint/models"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	repository[models.Organization]
	SearchOrganization(*models.Organization) ([]models.Organization, error)
}

type organizationRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.Organization]
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.Organization](db),
	}
}

func (r organizationRepositoryImpl) SearchOrganization(org *models.Organization) ([]models.Organization, error) {
	orgs := []models.Organization{}
	if err := r.db.Where("LOWER(name) LIKE LOWER(?)", org.Name+"%").Find(&orgs).Error; err != nil {
		return nil, err
	}

	return orgs, nil
}