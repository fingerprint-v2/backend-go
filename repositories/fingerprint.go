package repositories

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"gorm.io/gorm"
)

type FingerprintRepository interface {
	repository[models.Fingerprint, models.FingerprintFind]
	SearchFingerprint(req *dto.SearchFingerprintReq) (*[]models.Fingerprint, error)
}

type FingerprintRepositoryImpt struct {
	db *gorm.DB
	*repositoryImpl[models.Fingerprint, models.FingerprintFind]
}

func NewFingerprintRepository(db *gorm.DB) FingerprintRepository {
	return &FingerprintRepositoryImpt{
		db:             db,
		repositoryImpl: newRepository[models.Fingerprint, models.FingerprintFind](db),
	}
}

func (r *FingerprintRepositoryImpt) SearchFingerprint(req *dto.SearchFingerprintReq) (*[]models.Fingerprint, error) {
	fingerprints := new([]models.Fingerprint)

	//DB query
	if err := r.db.Find(fingerprints, req).Error; err != nil {
		return nil, err
	}

	return fingerprints, nil

}
