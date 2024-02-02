package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetUpMinioRouter(router fiber.Router, v dto.Validator, handler handlers.MinioHandler) {
	minio := router.Group("minio")
	minio.Post("/bucket/:bucket_name", handler.CreateBucket)
	minio.Post("/bucket/:bucket_name/model/:model_name", handler.UploadObject)
	minio.Get("/bucket/:bucket_name/model/:model_name", handler.DownloadObject)
}
