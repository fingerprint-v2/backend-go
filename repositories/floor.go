package repositories

import (
	"errors"

	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/utils"
	"gorm.io/gorm"
)

type FloorRepository interface {
	repository[models.Floor, models.FloorFind]
	SearchFloor(*dto.SearchFloorReq) (*[]models.Floor, error)
}

type floorRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.Floor, models.FloorFind]
}

func NewFloorRepository(db *gorm.DB) FloorRepository {
	return &floorRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.Floor, models.FloorFind](db),
	}
}

func (r *floorRepositoryImpl) SearchFloor(req *dto.SearchFloorReq) (*[]models.Floor, error) {
	floors := new([]models.Floor)

	// Construct the map from the request
	floorMap, err := utils.TypeConverter[map[string]interface{}](req)
	if err != nil {
		return nil, err
	}
	delete(*floorMap, "all")
	delete(*floorMap, "with_organization")
	delete(*floorMap, "with_site")
	delete(*floorMap, "with_building")
	delete(*floorMap, "with_point")

	// Make sure that map is not empty when "all" is false
	if len(*floorMap) == 0 && !req.All {
		return nil, errors.New("no valid search parameters provided")
	}
	if req.All {
		floorMap = &map[string]interface{}{}
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
	if req.WithPoint {
		db = db.Preload("Points") // plural
	}

	// DB query
	if err := db.Find(floors, *floorMap).Error; err != nil {
		return nil, err
	}
	return floors, nil
}
