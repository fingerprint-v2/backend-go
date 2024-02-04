package repositories

import (
	"github.com/fingerprint/models"
	"gorm.io/gorm"
)

type SiteRepository interface {
	repository[models.Site, models.SearchSite]
}

type siteRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.Site, models.SearchSite]
}

func NewSiteRepository(db *gorm.DB) SiteRepository {
	return &siteRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.Site, models.SearchSite](db),
	}
}
