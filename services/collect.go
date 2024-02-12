package services

import (
	"context"
	"fmt"

	"github.com/fingerprint/dto"
	"github.com/fingerprint/repositories"
)

type CollectService interface {
	Collect(ctx context.Context, req *dto.CreateSurveyReq) error
}

type collectServiceImpl struct {
	collectDeviceRepo repositories.CollectDeviceRepository
}

func NewCollectService(collectDeviceRepo repositories.CollectDeviceRepository) CollectService {
	return &collectServiceImpl{
		collectDeviceRepo: collectDeviceRepo,
	}
}

func (s *collectServiceImpl) Collect(ctx context.Context, req *dto.CreateSurveyReq) error {

	fmt.Println(req)
	return nil

}
