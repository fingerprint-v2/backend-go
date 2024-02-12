package services

import (
	"context"

	"github.com/fingerprint/repositories"
)

type CollectService interface {
	Collect(ctx context.Context) error
}

type collectServiceImpl struct {
	collectDeviceRepo repositories.CollectDeviceRepository
}

func NewCollectService(collectDeviceRepo repositories.CollectDeviceRepository) CollectService {
	return &collectServiceImpl{
		collectDeviceRepo: collectDeviceRepo,
	}
}

func (s *collectServiceImpl) Collect(ctx context.Context) error {
	return nil
}
