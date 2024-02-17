package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/minio/minio-go/v7"
)

type ObjectStorageService interface {
	CreateBucket(context.Context, string, minio.MakeBucketOptions) error
	UploadObject(context.Context, string, string, string, minio.PutObjectOptions) (*minio.UploadInfo, error)
	DownloadObject(context.Context, string, string, string, minio.GetObjectOptions) error
	WriteJSON(context.Context, string, string, interface{}) error
}

type objectStorageServiceImpl struct {
	client *minio.Client
}

func NewMinioService(client *minio.Client) ObjectStorageService {
	return &objectStorageServiceImpl{
		client: client,
	}
}
func (o *objectStorageServiceImpl) CreateBucket(ctx context.Context, bucketName string, opts minio.MakeBucketOptions) error {
	err := o.client.MakeBucket(ctx, bucketName, opts)
	if err != nil {
		return err
	}
	return nil
}

func (o *objectStorageServiceImpl) UploadObject(ctx context.Context, bucketName string, objectName string, path string, opts minio.PutObjectOptions) (*minio.UploadInfo, error) {
	uploadInfo, err := o.client.FPutObject(ctx, bucketName, objectName, path, opts)
	if err != nil {
		return nil, err
	}
	return &uploadInfo, nil
}

func (o *objectStorageServiceImpl) DownloadObject(ctx context.Context, bucketName string, objectName string, path string, opts minio.GetObjectOptions) error {
	if err := o.client.FGetObject(ctx, bucketName, objectName, path, opts); err != nil {
		return err
	}
	return nil
}

func (o *objectStorageServiceImpl) WriteJSON(ctx context.Context, bucketName string, objectName string, data interface{}) error {

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(dataBytes)

	// Upload the byte array with PutObject
	n, err := o.client.PutObject(ctx, bucketName, objectName, reader, int64(len(dataBytes)), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}
	fmt.Println("Successfully uploaded bytes: ", n)

	return nil
}
