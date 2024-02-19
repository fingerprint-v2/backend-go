package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

func GetMinioConfig() *MinioConfig {
	useSSL, err := strconv.ParseBool(os.Getenv("MINIO_USESSL"))
	if err != nil {
		return nil
	}

	return &MinioConfig{
		Endpoint:        os.Getenv("MINIO_HOST") + ":" + os.Getenv("MINIO_PORT"),
		AccessKeyID:     os.Getenv("MINIO_USER"),
		SecretAccessKey: os.Getenv("MINIO_PASSWORD"),
		UseSSL:          useSSL,
	}
}

func NewMinioClient() *minio.Client {

	configs := GetMinioConfig()

	// Initialize minio client object.
	minioClient, err := minio.New(configs.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(configs.AccessKeyID, configs.SecretAccessKey, ""),
		Secure: configs.UseSSL,
	})
	if err != nil {
		log.Fatalln(err.Error())
		return nil

	}

	return minioClient
}
