package repositories

import (
	"errors"

	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/utils"
	"gorm.io/gorm"
)

type SiteRepository interface {
	repository[models.Site, models.SiteFind]
	SearchSite(*dto.SearchSiteReq) (*[]models.Site, error)
}

type siteRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.Site, models.SiteFind]
}

func NewSiteRepository(db *gorm.DB) SiteRepository {
	return &siteRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.Site, models.SiteFind](db),
	}
}

func (r *siteRepositoryImpl) SearchSite(req *dto.SearchSiteReq) (*[]models.Site, error) {
	sites := new([]models.Site)

	// Construct the map from the request
	siteMap, err := utils.TypeConverter[map[string]interface{}](req)
	if err != nil {
		return nil, err
	}
	delete(*siteMap, "all")
	delete(*siteMap, "with_buildings")
	delete(*siteMap, "with_floors")
	delete(*siteMap, "with_points")

	// Make sure that map is not empty when "all" is false
	if len(*siteMap) == 0 && !req.All {
		return nil, errors.New("no valid search parameters provided")
	}
	if req.All {
		siteMap = &map[string]interface{}{}
	}

	// Optional preload
	db := r.db
	if req.WithBuildings {
		db = db.Preload("Buildings")
	}
	if req.WithFloors {
		db = db.Preload("Floors")
	}
	if req.WithPoints {
		db = db.Preload("Points")
	}

	// DB query
	if err := db.Find(sites, *siteMap).Error; err != nil {
		return nil, err
	}
	return sites, nil
}
