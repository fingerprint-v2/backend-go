package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fingerprint/ml"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCService interface {
	CheckModel() error
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
		log.Fatalf("%v", err)
	}
	fmt.Println(r)

	return nil
}
