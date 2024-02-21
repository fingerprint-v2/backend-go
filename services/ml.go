package services

import (
	"context"
	"fmt"

	"github.com/fingerprint/dto"
	"github.com/fingerprint/ml"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type MLService interface {
	CreateTraining(c context.Context, req *dto.CreateTrainingReq) error
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

func (s *mLServiceImpl) CreateTraining(c context.Context, req *dto.CreateTrainingReq) error {

	// Get points within the site
	points, err := s.pointRepo.GetPointsWithFingerprints(req.SiteID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Write file to cloud storage
	queueWriteJSON := &Job{Function: func() error {
		fmt.Println("Write JSON Started")
		// time.Sleep(3 * time.Second)
		if err := s.objectStorageService.WriteJSON(c, "training", "training.json", points); err != nil {
			return err
		}
		fmt.Println("Write JSON Finished")
		return nil
	}}

	// Train model
	grpcReq, err := utils.TypeConverter[[]*ml.Point](points)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	queueTrain := &Job{Function: func() error {
		fmt.Println("Train Started")
		_, err := s.gRPCService.Train(&ml.TrainReq{
			Name:   req.TrainingName,
			Points: *grpcReq,
		})

		if err != nil {
			return err
		}

		// time.Sleep(10 * time.Second)
		fmt.Println("Train Finished")
		return nil
	}}

	s.dispatcherService.Add(queueWriteJSON)
	s.dispatcherService.Add(queueTrain)
	// s.dispatcherService.Wait()

	return nil
}
