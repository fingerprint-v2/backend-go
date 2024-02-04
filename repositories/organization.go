package repositories

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrganizationRepository interface {
	repository[models.Organization, dto.SearchOrganizationReq]
	GetOrganizationsAllPreloads() (*[]models.Organization, error)
}

type organizationRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.Organization, dto.SearchOrganizationReq]
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return &organizationRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.Organization, dto.SearchOrganizationReq](db),
	}
}

func (r *organizationRepositoryImpl) GetOrganizationsAllPreloads() (*[]models.Organization, error) {
	organization := new([]models.Organization)
	if err := r.db.Preload(clause.Associations).Find(organization, map[string]interface{}{}).Error; err != nil {
		return nil, err
	}
	return organization, nil
}
