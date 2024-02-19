package configs

import (
	"log"

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

func NewGRPCConnection() (*grpc.ClientConn, func(), error) {
	conn, err := grpc.Dial(GetGRPCAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err.Error())
		return nil, nil, err
	}

	cleanup := func() {
		if err := conn.Close(); err != nil {
			log.Fatalln(err.Error())
		}

	}
	return conn, cleanup, nil
}

func NewGRPCClient(conn *grpc.ClientConn) ml.FingperintClient {
	client := ml.NewFingperintClient(conn)
	return client
}
