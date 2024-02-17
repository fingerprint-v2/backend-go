package configs

import (
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

func NewGRPCClient() *grpc.ClientConn {
	conn, err := grpc.Dial(GetGRPCAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil
	}
	return conn
}
