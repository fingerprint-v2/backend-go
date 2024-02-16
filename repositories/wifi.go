package repositories

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"gorm.io/gorm"
)

type WifiRepository interface {
	repository[models.Wifi, models.WifiFind]
	SearchWifi(*dto.SearchWifiReq) (*[]models.Wifi, error)
}

type wifiRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.Wifi, models.WifiFind]
}

func NewWifiRepository(db *gorm.DB) WifiRepository {
	return &wifiRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.Wifi, models.WifiFind](db),
	}
}

func (r *wifiRepositoryImpl) SearchWifi(req *dto.SearchWifiReq) (*[]models.Wifi, error) {
	wifis := new([]models.Wifi)
	// DB query
	if err := r.db.Find(wifis, req).Error; err != nil {
		return nil, err
	}
	return wifis, nil
}
