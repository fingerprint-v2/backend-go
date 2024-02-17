package handlers

import (
	"github.com/fingerprint/configs"
	"github.com/fingerprint/services"
	"github.com/fingerprint/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

type ObjectStorageHandler interface {
	CreateBucket(c *fiber.Ctx) error
	UploadObject(c *fiber.Ctx) error
	DownloadObject(c *fiber.Ctx) error
}

type objectStorageHandlerImpl struct {
	objectStorageService services.ObjectStorageService
}

func NewObjectStorageHandler(objectStorageService services.ObjectStorageService) ObjectStorageHandler {
	return &objectStorageHandlerImpl{
		objectStorageService: objectStorageService,
	}
}

// @Tags Minio
// @Summary Create Bucket
// @Description create Bucket
// @ID create-bucket
// @Accept json
// @Produce json
// @Param  bucket_name path string  true  "bucket name"
// @Success 200 {object} utils.ResponseSuccess[minio.UploadInfo]
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/minio/bucket/{bucket_name} [post]
func (h *objectStorageHandlerImpl) CreateBucket(c *fiber.Ctx) error {
	ctx := c.Context()
	bucketName := c.Params("bucket_name")
	if err := h.objectStorageService.CreateBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true}); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Create bucket sucessfully",
		Data:    nil,
	})
}

// @Tags Minio
// @Summary Upload Object
// @Description upload object
// @ID upload-object
// @Accept json
// @Produce json
// @Param  bucket_name path string  true  "bucket name"
// @Param  model_name path string  true  "model name"
// @Success 200 {object} utils.ResponseSuccess[minio.UploadInfo]
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/minio/bucket/{bucket_name}/model/{model_name} [post]
func (h *objectStorageHandlerImpl) UploadObject(c *fiber.Ctx) error {
	ctx := c.Context()
	bucketName := c.Params("bucket_name")
	modelName := c.Params("model_name")
	path := *configs.GetUploadPath() + modelName

	uploadInfo, err := h.objectStorageService.UploadObject(ctx, bucketName, modelName, path, minio.PutObjectOptions{})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[minio.UploadInfo]{
		Message: "Upload model sucessfully",
		Data:    *uploadInfo,
	})

}

// @Tags Minio
// @Summary Download Object
// @Description Download object
// @ID download-object
// @Accept json
// @Produce json
// @Param  bucket_name path string  true  "bucket name"
// @Param  model_name path string  true  "model name"
// @Success 200 {object} utils.ResponseSuccess[string]
// @Failure 500 {object} utils.ResponseError
// @Router /api/v1/minio/bucket/{bucket_name}/model/{model_name} [get]
func (h *objectStorageHandlerImpl) DownloadObject(c *fiber.Ctx) error {
	ctx := c.Context()
	bucketName := c.Params("bucket_name")
	modelName := c.Params("model_name")
	dest := *configs.GetDownloadPath() + modelName

	if err := h.objectStorageService.DownloadObject(ctx, bucketName, modelName, dest, minio.GetObjectOptions{}); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(utils.ResponseSuccess[interface{}]{
		Message: "Download model sucessfully",
		Data:    nil,
	})
}
