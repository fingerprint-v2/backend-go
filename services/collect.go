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
	wifiRepo          repositories.WifiRepository
	pointRepo         repositories.PointRepository
	siteRepo          repositories.SiteRepository
}

func NewCollectService(
	collectDeviceRepo repositories.CollectDeviceRepository,
	uploadrepo repositories.UploadRepository,
	fingerprintRepo repositories.FingerprintRepository,
	wifiRepo repositories.WifiRepository,
	pointRepo repositories.PointRepository,
	siteRepo repositories.SiteRepository,
) CollectService {
	return &collectServiceImpl{
		collectDeviceRepo: collectDeviceRepo,
		uploadRepo:        uploadrepo,
		fingerprintRepo:   fingerprintRepo,
		wifiRepo:          wifiRepo,
		pointRepo:         pointRepo,
		siteRepo:          siteRepo,
	}
}

func (s *collectServiceImpl) Collect(req *dto.CreateSurveyReq, user *models.User) error {

	// Collect Device
	collectDeviceReq, err := utils.TypeConverter[models.CollectDevice](req.CollectDevice)
	if err != nil {
		return err
	}

	var collectDeviceID string
	collectDevices, err := s.collectDeviceRepo.Find(&models.CollectDeviceFind{DeviceUID: collectDeviceReq.DeviceUID})
	if err != nil {
		return err
	}
	if len(*collectDevices) > 0 {
		// Update
		collectDevice := (*collectDevices)[0]
		if err := s.collectDeviceRepo.Update(collectDevice.ID.String(), collectDeviceReq); err != nil {
			return err
		}
		collectDeviceID = collectDevice.ID.String()
	} else {
		// Create
		if err := s.collectDeviceRepo.Create(collectDeviceReq); err != nil {
			return err
		}
		collectDeviceID = collectDeviceReq.ID.String()
	}

	// Upload
	upload := &models.Upload{
		UserID:       user.ID.String(),
		ScanMode:     req.ScanMode,
		ScanInterval: req.ScanInterval,
	}
	if err := s.uploadRepo.Create(upload); err != nil {
		return err
	}
	uploadID := upload.ID.String()

	// Check Valid Collect Mode (just to make sure)
	if err := s.CheckValidCollectMode(req.Mode); err != nil {
		return err
	}
	mode := req.Mode

	// Determine SiteID
	var siteID string
	if mode == constants.SUPERVISED.String() {
		point, err := s.pointRepo.Get(req.PointLabelID)
		if err != nil {
			return err
		}
		siteID = point.SiteID
	} else if (mode == constants.UNSUPERVISED.String()) || (mode == constants.PREDICTION.String()) {
		if req.SiteID == "" {
			return errors.New("site_id is required")
		}
		siteID = req.SiteID
	}

	// Determine OrganizationID
	site, err := s.siteRepo.Get(siteID)
	if err != nil {
		return err
	}
	organizationID := site.OrganizationID

	// Determine PointID
	pointLabelID := new(string)
	if mode == constants.SUPERVISED.String() {
		point, err := s.pointRepo.Get(req.PointLabelID)
		if err != nil {
			return err
		}
		tempStr := point.ID.String()
		pointLabelID = &tempStr
	}

	// Create Fingerprint
	fingerprints := []*models.Fingerprint{}
	for _, fingerprintReq := range req.Fingerprints {
		wifis := []models.Wifi{}
		for _, wifiReq := range fingerprintReq.Wifis {
			wifi, err := utils.TypeConverter[models.Wifi](wifiReq)
			if err != nil {
				return err
			}
			wifis = append(wifis, *wifi)
		}
		fingerprint := &models.Fingerprint{
			Mode:            req.Mode,
			CollectDeviceID: collectDeviceID,
			SiteID:          siteID,
			OrganizationID:  organizationID,
			UploadID:        uploadID,
			Wifis:           wifis,
			PointLabelID:    pointLabelID,
		}
		fingerprints = append(fingerprints, fingerprint)
	}

	if err := s.fingerprintRepo.CreateMultiple(fingerprints); err != nil {
		return err
	}

	// Create Wifis

	// wifis := req.Fingerprints

	fmt.Println(req, user)

	return nil

}

func (s *collectServiceImpl) CheckValidCollectMode(mode string) error {

	modes := []constants.CollectMode{constants.PREDICTION, constants.SUPERVISED, constants.UNSUPERVISED}
	for _, value := range modes {
		if mode == value.String() {
			return nil
		}
	}
	return errors.New("invalid role")
}
