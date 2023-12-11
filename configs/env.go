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
		return nil
	}

	return []byte(refreshTokenSignature)
}
