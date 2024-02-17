package routers

import (
	"github.com/fingerprint/dto"
	"github.com/fingerprint/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetUpObjectStorageRouter(router fiber.Router, v dto.Validator, handler handlers.ObjectStorageHandler) {
	minio := router.Group("files")
	minio.Post("/bucket/:bucket_name", handler.CreateBucket)
	minio.Post("/bucket/:bucket_name/model/:model_name", handler.UploadObject)
	minio.Get("/bucket/:bucket_name/model/:model_name", handler.DownloadObject)
}
