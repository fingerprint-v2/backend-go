package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fingerprint/ml"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCService interface {
	CheckModel() error
	Train(*ml.TrainReq) (*ml.TrainRes, error)
}

type gRPCServiceImpl struct {
	client ml.FingperintClient
}

func NewGRPCService(client ml.FingperintClient) GRPCService {
	return &gRPCServiceImpl{
		client: client,
	}
}

func (s *gRPCServiceImpl) CheckModel() error {

	grpcCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := s.client.LoadModel(grpcCtx, &ml.LoadModelReq{
		Path: "model",
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	r, err := s.client.CheckModel(grpcCtx, &emptypb.Empty{})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	fmt.Println(r)

	return nil
}

func (s *gRPCServiceImpl) Train(req *ml.TrainReq) (*ml.TrainRes, error) {
	grpcCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := s.client.Train(grpcCtx, req)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())

	}
	return r, nil
}
