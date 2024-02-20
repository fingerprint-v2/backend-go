package services

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/ml"
	"github.com/fingerprint/repositories"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
)

type MLService interface {
	CreateTraining(c *fiber.Ctx, req *dto.CreateTrainingReq) (*ml.TrainRes, error)
}

type mLServiceImpl struct {
	pointRepo            repositories.PointRepository
	objectStorageService ObjectStorageService
	gRPCService          GRPCService
}

func NewMLService(
	objectStrorageService ObjectStorageService,
	gRPCService GRPCService,
	pointRepo repositories.PointRepository,
) MLService {
	return &mLServiceImpl{
		objectStorageService: objectStrorageService,
		gRPCService:          gRPCService,
		pointRepo:            pointRepo,
	}
}

func (s *mLServiceImpl) CreateTraining(c *fiber.Ctx, req *dto.CreateTrainingReq) (*ml.TrainRes, error) {

	// Get points within the site
	points, err := s.pointRepo.GetPointsWithFingerprints(req.SiteID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// fmt.Println(points)

	if err := s.objectStorageService.WriteJSON(c.Context(), "training", "training.json", points); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// grpcReq := &ml.TrainReq{
	// 	Name:   "Model1",
	// 	Points: []*ml.Point{},
	// }

	grpcReq, err := utils.TypeConverter[[]*ml.Point](points)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	r, err := s.gRPCService.Train(&ml.TrainReq{
		Name:   req.TrainingName,
		Points: *grpcReq,
	})

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return r, nil
}
