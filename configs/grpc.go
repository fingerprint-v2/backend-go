package configs

import (
	"github.com/fingerprint/ml"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCConfig struct {
	Host string
	Port string
}

func GetGRPCConfig() *GRPCConfig {
	return &GRPCConfig{
		Host: "localhost",
		Port: "50051",
	}
}

func GetGRPCAddress() string {
	configs := GetGRPCConfig()
	return configs.Host + ":" + configs.Port
}

func NewGRPCConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(GetGRPCAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil
	}
	return conn, nil
}

func NewGRPCClient(conn *grpc.ClientConn) ml.FingperintClient {
	client := ml.NewFingperintClient(conn)
	return client
}
