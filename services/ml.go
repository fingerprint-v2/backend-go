package services

import (
	"fmt"
	"time"

	"github.com/fingerprint/dto"
	"github.com/fingerprint/repositories"
	"github.com/gofiber/fiber/v2"
)

type MLService interface {
	CreateTraining(c *fiber.Ctx, req *dto.CreateTrainingReq) error
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

func (s *mLServiceImpl) CreateTraining(c *fiber.Ctx, req *dto.CreateTrainingReq) error {

	// Get points within the site
	// points, err := s.pointRepo.GetPointsWithFingerprints(req.SiteID)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }
	// fmt.Println(points)

	// if err := s.objectStorageService.WriteJSON(c.Context(), "training", "training.json", points); err != nil {
	// 	return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }

	writeJSON := &Job{Function: func() error {
		fmt.Println("Write JSON Started")
		time.Sleep(3 * time.Second)
		// if err := s.objectStorageService.WriteJSON(c.Context(), "training", "training.json", points); err != nil {
		// 	return err
		// }
		fmt.Println("Write JSON Finished")
		return nil
	}}

	// grpcReq := &ml.TrainReq{
	// 	Name:   "Model1",
	// 	Points: []*ml.Point{},
	// }

	// grpcReq, err := utils.TypeConverter[[]*ml.Point](points)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }

	// r, err := s.gRPCService.Train(&ml.TrainReq{
	// 	Name:   req.TrainingName,
	// 	Points: *grpcReq,
	// })

	// if err != nil {
	// 	return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }

	// var r *ml.TrainRes
	train := &Job{Function: func() error {
		fmt.Println("Train Started")
		// _, err := s.gRPCService.Train(&ml.TrainReq{
		// 	Name:   req.TrainingName,
		// 	Points: *grpcReq,
		// })

		// if err != nil {
		// 	return err
		// }

		// r = result
		time.Sleep(3 * time.Second)
		// fmt.Println(r)
		fmt.Println("Train Finished")
		return nil
	}}

	s.dispatcherService.Add(writeJSON)
	s.dispatcherService.Add(train)
	s.dispatcherService.Wait()

	return nil
}
