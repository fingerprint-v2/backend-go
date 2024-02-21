package services

import (
	"context"

	"github.com/fingerprint/dto"
	"github.com/fingerprint/ml"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type MLService interface {
	CreateTraining(ctx context.Context, req *dto.CreateTrainingReq) (*ml.TrainRes, error)
}

type mLServiceImpl struct {
	pointRepo            repositories.PointRepository
	objectStorageService ObjectStorageService
	gRPCService          GRPCService
	dispatcherService    DispatcherService
}

func NewMLService(
	objectStrorageService ObjectStorageService,
	gRPCService GRPCService,
	pointRepo repositories.PointRepository,
	dispatcherService DispatcherService,
) MLService {
	return &mLServiceImpl{
		objectStorageService: objectStrorageService,
		gRPCService:          gRPCService,
		pointRepo:            pointRepo,
		dispatcherService:    dispatcherService,
	}
}
func (s *mLServiceImpl) CreateTraining(ctx context.Context, req *dto.CreateTrainingReq) (*ml.TrainRes, error) {
	s.dispatcherService.Start(ctx)

	points, err := s.pointRepo.GetPointsWithFingerprints(req.SiteID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	writeJSON := &Job{Function: func() error {
		if err := s.objectStorageService.WriteJSON(ctx, "training", "training.json", points); err != nil {
			return err
		}
		return nil
	}}

	grpcReq, err := utils.TypeConverter[[]*ml.Point](points)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	var r *ml.TrainRes
	train := &Job{Function: func() error {
		result, err := s.gRPCService.Train(&ml.TrainReq{
			Name:   req.TrainingName,
			Points: *grpcReq,
		})

		if err != nil {
			return err
		}

		r = result
		return nil
	}}

	s.dispatcherService.Add(writeJSON)
	s.dispatcherService.Add(train)
	s.dispatcherService.Wait()

	return r, nil
}
