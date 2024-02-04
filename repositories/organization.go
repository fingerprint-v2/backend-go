package repositories

import (
	"github.com/fingerprint/models"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	repository[models.Organization, models.OrganizationFind]
	SearchOrganization() (*[]models.Organization, error)
}

type organizationRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.Organization, models.OrganizationFind]
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.Organization, models.OrganizationFind](db),
	}
}

func (r *organizationRepositoryImpl) SearchOrganization() (*[]models.Organization, error) {
	organization := new([]models.Organization)
	// if err := r.db.Preload(clause.Associations).Find(organization, map[string]interface{}{}).Error; err != nil {
	// 	return nil, err
	// }

	db := r.db

	db = db.Preload("Users")
	db = db.Preload("Sites")

	if err := db.Find(organization, map[string]interface{}{}).Error; err != nil {
		return nil, err
	}

	return organization, nil
}
