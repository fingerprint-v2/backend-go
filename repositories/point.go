package repositories

import (
	"errors"

	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/utils"
	"gorm.io/gorm"
)

type PointRepository interface {
	repository[models.Point, models.PointFind]
	SearchPoint(*dto.SearchPointReq) (*[]models.Point, error)
	GetPointsWithFingerprints(siteID string) (*[]models.Point, error)
}

type pointRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.Point, models.PointFind]
}

func NewPointRepository(db *gorm.DB) PointRepository {
	return &pointRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.Point, models.PointFind](db),
	}
}

func (r *pointRepositoryImpl) SearchPoint(req *dto.SearchPointReq) (*[]models.Point, error) {
	points := new([]models.Point)

	// Construct the map from the request
	pointMap, err := utils.TypeConverter[map[string]interface{}](req)
	if err != nil {
		return nil, err
	}
	delete(*pointMap, "all")
	delete(*pointMap, "with_organization")
	delete(*pointMap, "with_site")
	delete(*pointMap, "with_building")
	delete(*pointMap, "with_floor")
	delete(*pointMap, "with_member")
	delete(*pointMap, "with_fingerprint")

	// Make sure that map is not empty when "all" is false
	if len(*pointMap) == 0 && !req.All {
		return nil, errors.New("no valid search parameters provided")
	}
	if req.All {
		pointMap = &map[string]interface{}{}
	}

	// Optional preload
	db := r.db
	if req.WithOrganization {
		db = db.Preload("Organization") // singular
	}
	if req.WithSite {
		db = db.Preload("Site") // singular
	}
	if req.WithBuilding {
		db = db.Preload("Building") // singular
	}
	if req.WithFloor {
		db = db.Preload("Floor") // singular
	}
	if req.WithMember {
		db = db.Preload("Points") // plural
	}
	if req.WithFingerprint {
		db = db.Preload("Fingerprints") // plural
	}

	// DB query
	if err := db.Find(points, *pointMap).Error; err != nil {
		return nil, err
	}
	return points, nil
}

func (r *pointRepositoryImpl) GetPointsWithFingerprints(siteID string) (*[]models.Point, error) {
	points := new([]models.Point)

	if err := r.db.Debug().Preload("Fingerprints.Wifis").Find(points, "site_id = ?", siteID).Error; err != nil {
		return nil, err
	}
	return points, nil
}
