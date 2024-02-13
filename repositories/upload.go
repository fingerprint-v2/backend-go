package repositories

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"gorm.io/gorm"
)

type UploadRepository interface {
	repository[models.Upload, models.UploadFind]
	SearchUpload(*dto.SearchUploadReq) (*[]models.Upload, error)
}

type uploadRepositoryImpl struct {
	db *gorm.DB
	*repositoryImpl[models.Upload, models.UploadFind]
}

func NewUploadRepository(db *gorm.DB) UploadRepository {
	return &uploadRepositoryImpl{
		db:             db,
		repositoryImpl: newRepository[models.Upload, models.UploadFind](db),
	}
}

func (r *uploadRepositoryImpl) SearchUpload(req *dto.SearchUploadReq) (*[]models.Upload, error) {
	uploads := new([]models.Upload)

	// DB query
	if err := r.db.Find(uploads, req).Error; err != nil {
		return nil, err
	}
	return uploads, nil
}
