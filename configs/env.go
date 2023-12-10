package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
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
func GetEncryptKey() string {
	key, ok := os.LookupEnv("COOKIE_ENCRYPT_KEY")
	if !ok {
		return encryptcookie.GenerateKey()
	}
	return key
}

func GetJWTSignature() (string, error) {
	jwtSignature, ok := os.LookupEnv("JWT_SIGNATURE")
	if !ok || jwtSignature == "" {
		return "", fmt.Errorf("environment variable JWT_SIGNATURE is not set or empty")
	}

	return jwtSignature, nil
}
