package repositories

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"gorm.io/gorm"
)

type CollectDeviceRepository interface {
	repository[models.CollectDevice, models.CollectDeviceFind]
	SearchCollectDevice(*dto.SearchCollectDeviceReq) (*[]models.CollectDevice, error)
}

type collectDeviceRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.CollectDevice, models.CollectDeviceFind]
}

func NewCollectDeviceRepository(db *gorm.DB) CollectDeviceRepository {
	return &collectDeviceRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.CollectDevice, models.CollectDeviceFind](db),
	}
}

func (r *collectDeviceRepositoryImpl) SearchCollectDevice(req *dto.SearchCollectDeviceReq) (*[]models.CollectDevice, error) {
	collectDevices := new([]models.CollectDevice)
	return collectDevices, nil
}
