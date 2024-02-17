package services

import (
	"github.com/fingerprint/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

type TrainingService interface {
	CreateTraining(c *fiber.Ctx, req *dto.CreateTrainingReq) error
}

type trainingServiceImpl struct {
	objectStorageService ObjectStorageService
	gRPCService          GRPCService
}

func NewTrainingService(objectStrorageService ObjectStorageService, gRPCService GRPCService) TrainingService {
	return &trainingServiceImpl{
		objectStorageService: objectStrorageService,
		gRPCService:          gRPCService,
	}
}

func (s *trainingServiceImpl) CreateTraining(c *fiber.Ctx, req *dto.CreateTrainingReq) error {

	if err := s.objectStorageService.CreateBucket(c.Context(), "training", minio.MakeBucketOptions{
		Region:        "us-east-1",
		ObjectLocking: false,
	}); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := s.objectStorageService.WriteJSON(c.Context(), "training", "training.json", req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err := s.gRPCService.NewTodo(); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
