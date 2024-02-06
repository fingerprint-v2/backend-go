package repositories

import (
	"errors"

	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/utils"
	"gorm.io/gorm"
)

type BuildingRepository interface {
	repository[models.Building, models.BuildingFind]
	SearchBuilding(*dto.SearchBuildingReq) (*[]models.Building, error)
}

type buildingRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.Building, models.BuildingFind]
}

func NewBuildingRepository(db *gorm.DB) BuildingRepository {
	return &buildingRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.Building, models.BuildingFind](db),
	}
}

func (r *buildingRepositoryImpl) SearchBuilding(req *dto.SearchBuildingReq) (*[]models.Building, error) {
	buildings := new([]models.Building)

	// Construct the map from the request
	buildingMap, err := utils.TypeConverter[map[string]interface{}](req)
	if err != nil {
		return nil, err
	}
	delete(*buildingMap, "all")
	delete(*buildingMap, "with_organization")
	delete(*buildingMap, "with_site")
	delete(*buildingMap, "with_floor")
	delete(*buildingMap, "with_point")

	// Make sure that map is not empty when "all" is false
	if len(*buildingMap) == 0 && !req.All {
		return nil, errors.New("no valid search parameters provided")
	}
	if req.All {
		buildingMap = &map[string]interface{}{}
	}

	// Optional preload
	db := r.db
	if req.WithOrganization {
		db = db.Preload("Organization") // singular
	}
	if req.WithSite {
		db = db.Preload("Site") // singular
	}
	if req.WithFloor {
		db = db.Preload("Floors") // plural
	}
	if req.WithPoint {
		db = db.Preload("Points") // plural
	}

	// DB query
	if err := db.Find(buildings, *buildingMap).Error; err != nil {
		return nil, err
	}
	return buildings, nil
}
