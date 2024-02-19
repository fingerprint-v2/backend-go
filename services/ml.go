package services

import (
	"fmt"

	"github.com/fingerprint/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

type MLService interface {
	CreateTraining(c *fiber.Ctx, req *dto.CreateTrainingReq) error
}

type mLServiceImpl struct {
	objectStorageService ObjectStorageService
	gRPCService          GRPCService
}

func NewMLService(objectStrorageService ObjectStorageService, gRPCService GRPCService) MLService {
	return &mLServiceImpl{
		objectStorageService: objectStrorageService,
		gRPCService:          gRPCService,
	}
}

func (s *mLServiceImpl) CreateTraining(c *fiber.Ctx, req *dto.CreateTrainingReq) error {

	if err := s.objectStorageService.CreateBucket(c.Context(), "training", minio.MakeBucketOptions{
		Region:        "us-east-1",
		ObjectLocking: false,
	}); err != nil {
		// return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		fmt.Println(err.Error())
	}

	if err := s.objectStorageService.WriteJSON(c.Context(), "training", "training.json", req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := s.gRPCService.CheckModel(); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
