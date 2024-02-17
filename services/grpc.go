package services

import (
	todo "github.com/fingerprint/todo"
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

	c := todo.NewTodoServiceClient(s.client)

	// // Contact the server and print out its response.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// r, err := c.GetTodos(ctx, &pb.Empty{})
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }
	// fmt.Println(r)

	return nil
}
