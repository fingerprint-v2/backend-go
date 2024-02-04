package repositories

import (
	"errors"

	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/utils"
	"gorm.io/gorm"
)

type OrganizationRepository interface {
	repository[models.Organization, models.OrganizationFind]
	SearchOrganization(*dto.SearchOrganizationReq) (*[]models.Organization, error)
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

func (r *organizationRepositoryImpl) SearchOrganization(req *dto.SearchOrganizationReq) (*[]models.Organization, error) {
	orgs := new([]models.Organization)

	// Construct the map from the request
	orgMap, err := utils.TypeConverter[map[string]interface{}](req)
	if err != nil {
		return nil, err
	}
	delete(*orgMap, "with_sites")
	delete(*orgMap, "with_users")
	delete(*orgMap, "all")
	// Make sure that map is not empty when "all" is false
	if len(*orgMap) == 0 && !req.All {
		return nil, errors.New("no valid search parameters provided")
	}
	if req.All {
		orgMap = &map[string]interface{}{}
	}

	// Optional preload
	db := r.db
	if req.WithSites {
		db = db.Preload("Sites")
	}
	if req.WithUsers {
		db = db.Preload("Users")
	}

	// DB query
	if err := db.Find(orgs, *orgMap).Error; err != nil {
		return nil, err
	}
	return orgs, nil
}
