package repositories

import (
	"github.com/fingerprint/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrganizationRepository interface {
	repository[models.Organization, models.SearchOrganization]
	GetOrganizationsPreloads() (*[]models.Organization, error)
}

type organizationRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.Organization, models.SearchOrganization]
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.Organization, models.SearchOrganization](db),
	}
}

func (r *organizationRepositoryImpl) GetOrganizationsPreloads() (*[]models.Organization, error) {
	organization := new([]models.Organization)
	if err := r.db.Preload(clause.Associations).Find(organization, map[string]interface{}{}).Error; err != nil {
		return nil, err
	}
	return organization, nil
}
