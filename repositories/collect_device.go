package repositories

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CollectDeviceRepository interface {
	repository[models.CollectDevice, models.CollectDeviceFind]
	SearchCollectDevice(*dto.SearchCollectDeviceReq) (*[]models.CollectDevice, error)
	CreateOrUpdateCollectDevice(req *dto.CreateCollectDeviceReq) (string, error)
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

	// DB query
	if err := r.db.Find(collectDevices, req).Error; err != nil {
		return nil, err
	}

	return collectDevices, nil
}

func (r *collectDeviceRepositoryImpl) CreateOrUpdateCollectDevice(req *dto.CreateCollectDeviceReq) (string, error) {

	var collectDeviceID string

	newCollectDevice, err := utils.TypeConverter[models.CollectDevice](req)
	if err != nil {
		return "", err
	}
	collectDevices, err := r.Find(&models.CollectDeviceFind{DeviceUID: req.DeviceUID})
	if err != nil {
		return "", err
	}
	if len(*collectDevices) == 1 {
		// Update
		existingCollectDevice := (*collectDevices)[0]
		if err := r.Update(existingCollectDevice.ID.String(), newCollectDevice); err != nil {
			return "", err
		}
		collectDeviceID = existingCollectDevice.ID.String()
	} else if len(*collectDevices) == 0 {
		// Create
		if err := r.Create(newCollectDevice); err != nil {
			return "", err
		}
		collectDeviceID = newCollectDevice.ID.String()
	} else {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Duplicated device UID")

	}
	return collectDeviceID, nil
}
