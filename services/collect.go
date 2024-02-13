package services

import (
	"errors"
	"fmt"

	"github.com/fingerprint/constants"
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/utils"
)

type CollectService interface {
	Collect(req *dto.CreateSurveyReq, user *models.User) error
}

type collectServiceImpl struct {
	collectDeviceRepo repositories.CollectDeviceRepository
	uploadRepo        repositories.UploadRepository
	fingerprintRepo   repositories.FingerprintRepository
}

func NewCollectService(
	collectDeviceRepo repositories.CollectDeviceRepository,
	uploadrepo repositories.UploadRepository,
	fingerprintRepo repositories.FingerprintRepository,
) CollectService {
	return &collectServiceImpl{
		collectDeviceRepo: collectDeviceRepo,
		uploadRepo:        uploadrepo,
		fingerprintRepo:   fingerprintRepo,
	}
}

func (s *collectServiceImpl) Collect(req *dto.CreateSurveyReq, user *models.User) error {

	// Collect Device
	collectDeviceReq, err := utils.TypeConverter[models.CollectDevice](req.Device)
	if err != nil {
		return err
	}

	devices, err := s.collectDeviceRepo.Find(&models.CollectDeviceFind{DeviceUID: collectDeviceReq.DeviceUID})
	if err != nil {
		return err
	}
	if len(*devices) > 0 {
		// Update
		device := (*devices)[0]
		if err := s.collectDeviceRepo.Update(device.ID.String(), collectDeviceReq); err != nil {
			return err
		}
	} else {
		// Create
		if err := s.collectDeviceRepo.Create(collectDeviceReq); err != nil {
			return err
		}
	}

	// Upload
	upload := &models.Upload{
		UserID: user.ID.String(),
	}
	if err := s.uploadRepo.Create(upload); err != nil {
		return err
	}

	// Fingerprint

	fmt.Println(req, user)
	return nil

}

func (s *authServiceImpl) CheckValidCollectMode(mode string) error {

	modes := []constants.CollectMode{constants.PREDICTION, constants.SUPERVISED, constants.UNSUPERVISED}
	for _, value := range modes {
		if mode == value.String() {
			return nil
		}
	}
	return errors.New("invalid role")
}
