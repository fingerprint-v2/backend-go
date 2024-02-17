package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fingerprint/grpcmicro"
	"google.golang.org/grpc"
)

type GRPCService interface {
	NewTodo() error
}

type gRPCServiceImpl struct {
	client *grpc.ClientConn
}

func NewGRPCService(client *grpc.ClientConn) GRPCService {
	return &gRPCServiceImpl{
		client: client,
	}
}

func (s *gRPCServiceImpl) NewTodo() error {

	c := grpcmicro.NewTodoServiceClient(s.client)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetTodos(ctx, &grpcmicro.Empty{})
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(r)

	return nil
}
