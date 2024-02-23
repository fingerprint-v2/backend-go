package services

import (
	"strings"

	"github.com/fingerprint/constants"
	"github.com/fingerprint/dto"
	"github.com/fingerprint/models"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/utils"
)

type LocalizeService interface {
	CreateSupervisedSurvey(req *dto.CreateSupervisedSurveyReq, user *models.User) error
	CreateUnsupervisedSurvey(req *dto.CreateUnsupervisedSurveyReq, user *models.User) error
}

type localizeServiceImpl struct {
	collectDeviceRepo repositories.CollectDeviceRepository
	uploadRepo        repositories.UploadRepository
	fingerprintRepo   repositories.FingerprintRepository
	wifiRepo          repositories.WifiRepository
	pointRepo         repositories.PointRepository
	siteRepo          repositories.SiteRepository
}

func NewLocalizeService(
	collectDeviceRepo repositories.CollectDeviceRepository,
	uploadrepo repositories.UploadRepository,
	fingerprintRepo repositories.FingerprintRepository,
	wifiRepo repositories.WifiRepository,
	pointRepo repositories.PointRepository,
	siteRepo repositories.SiteRepository,
) LocalizeService {
	return &localizeServiceImpl{
		collectDeviceRepo: collectDeviceRepo,
		uploadRepo:        uploadrepo,
		fingerprintRepo:   fingerprintRepo,
		wifiRepo:          wifiRepo,
		pointRepo:         pointRepo,
		siteRepo:          siteRepo,
	}
}

func (s *localizeServiceImpl) CreateSupervisedSurvey(req *dto.CreateSupervisedSurveyReq, user *models.User) error {

	// Collect Device
	collectDeviceReq, err := utils.TypeConverter[dto.CreateCollectDeviceReq](req.CollectDevice)
	if err != nil {
		return err
	}
	collectDeviceID, err := s.collectDeviceRepo.CreateOrUpdateCollectDevice(collectDeviceReq)
	if err != nil {
		return err
	}

	// Upload
	upload := &models.Upload{
		UserID:       user.ID.String(),
		UploadMode:   constants.SURVEY_SUPERVISED.String(),
		ScanMode:     req.ScanMode,
		ScanInterval: req.ScanInterval,
	}
	if err := s.uploadRepo.Create(upload); err != nil {
		return err
	}
	uploadID := upload.ID.String()

	// Determine PointLabelID
	point, err := s.pointRepo.Get(req.PointLabelID)
	if err != nil {
		return err
	}
	pointLabelID := point.ID.String()

	// Create Fingerprint
	fingerprints := []*models.Fingerprint{}
	for _, fingerprintReq := range req.Fingerprints {
		wifis := []models.Wifi{}
		for _, wifiReq := range fingerprintReq.Wifis {
			wifi, err := utils.TypeConverter[models.Wifi](wifiReq)
			if err != nil {
				return err
			}
			wifi.BSSID = strings.ToLower(wifi.BSSID)
			wifis = append(wifis, *wifi)
		}
		fingerprint := &models.Fingerprint{
			CollectDeviceID: collectDeviceID,
			SiteID:          point.SiteID,
			OrganizationID:  point.OrganizationID,
			UploadID:        uploadID,
			Wifis:           wifis,
			PointLabelID:    &pointLabelID,
			IsCurrent:       true,
		}
		fingerprints = append(fingerprints, fingerprint)
	}

	if err := s.fingerprintRepo.CreateMultiple(fingerprints); err != nil {
		return err
	}

	return nil

}

func (s *localizeServiceImpl) CreateUnsupervisedSurvey(req *dto.CreateUnsupervisedSurveyReq, user *models.User) error {

	// Collect Device
	collectDeviceReq, err := utils.TypeConverter[dto.CreateCollectDeviceReq](req.CollectDevice)
	if err != nil {
		return err
	}
	collectDeviceID, err := s.collectDeviceRepo.CreateOrUpdateCollectDevice(collectDeviceReq)
	if err != nil {
		return err
	}

	// Upload
	upload := &models.Upload{
		UserID:       user.ID.String(),
		UploadMode:   constants.SURVEY_UNSUPERVISED.String(),
		ScanMode:     req.ScanMode,
		ScanInterval: req.ScanInterval,
	}
	if err := s.uploadRepo.Create(upload); err != nil {
		return err
	}
	uploadID := upload.ID.String()

	// Determine SiteID
	site, err := s.siteRepo.Get(req.SiteID)
	if err != nil {
		return err
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
			wifi.BSSID = strings.ToLower(wifi.BSSID)
			wifis = append(wifis, *wifi)
		}
		fingerprint := &models.Fingerprint{
			CollectDeviceID: collectDeviceID,
			SiteID:          site.ID.String(),
			OrganizationID:  site.OrganizationID,
			UploadID:        uploadID,
			Wifis:           wifis,
			PointLabelID:    nil,
			IsCurrent:       true,
		}
		fingerprints = append(fingerprints, fingerprint)
	}

	if err := s.fingerprintRepo.CreateMultiple(fingerprints); err != nil {
		return err
	}

	return nil
}
