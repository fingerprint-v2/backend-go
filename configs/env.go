package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitialEnv(path string) {
	if err := godotenv.Load(path); err != nil {
		log.Println("No .env file found")
	}
}

func GetPort() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		return "8000"
	}
	return port
}

func GetAccessTokenSignature() []byte {
	accessTokenSignature, ok := os.LookupEnv("JWT_ACCESS_SIGNATURE")
	if !ok || accessTokenSignature == "" {
		return nil
	}

	return []byte(accessTokenSignature)
}

func GetRefreshTokenSignature() []byte {
	refreshTokenSignature, ok := os.LookupEnv("JWT_REFRESH_SIGNATURE")
	if !ok || refreshTokenSignature == "" {
		log.Println("JWT_REFRESH_SIGNATURE not set")
		return nil
	}

	return []byte(refreshTokenSignature)
}

func GetUploadPath() *string {
	uploadPath, ok := os.LookupEnv("MINIO_UPLOAD_PATH")
	if !ok || uploadPath == "" {
		log.Println("MINIO_UPLOAD_PATH not set")
		return nil
	}
	return &uploadPath
}

func GetDownloadPath() *string {
	uploadPath, ok := os.LookupEnv("MINIO_DOWNLOAD_PATH")
	if !ok || uploadPath == "" {
		log.Println("MINIO_DOWNLOAD_PATH not set")
		return nil
	}
	return &uploadPath
}
