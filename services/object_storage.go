package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/minio/minio-go/v7"
)

type ObjectStorageService interface {
	CreateBucket(ctx context.Context, bucketName string, opts minio.MakeBucketOptions) error
	UploadObject(context.Context, string, string, string, minio.PutObjectOptions) (*minio.UploadInfo, error)
	DownloadObject(context.Context, string, string, string, minio.GetObjectOptions) error
	WriteJSON(ctx context.Context, bucketName string, objectName string, data interface{}) error
}

type objectStorageServiceImpl struct {
	client *minio.Client
}

func NewObjectStorageService(client *minio.Client) ObjectStorageService {
	return &objectStorageServiceImpl{
		client: client,
	}
}
func (s *objectStorageServiceImpl) CreateBucket(ctx context.Context, bucketName string, opts minio.MakeBucketOptions) error {
	err := s.client.MakeBucket(ctx, bucketName, opts)
	if err != nil {
		return err
	}
	return nil
}

func (s *objectStorageServiceImpl) UploadObject(ctx context.Context, bucketName string, objectName string, path string, opts minio.PutObjectOptions) (*minio.UploadInfo, error) {
	uploadInfo, err := s.client.FPutObject(ctx, bucketName, objectName, path, opts)
	if err != nil {
		return nil, err
	}
	return &uploadInfo, nil
}

func (s *objectStorageServiceImpl) DownloadObject(ctx context.Context, bucketName string, objectName string, path string, opts minio.GetObjectOptions) error {
	if err := s.client.FGetObject(ctx, bucketName, objectName, path, opts); err != nil {
		return err
	}
	return nil
}

func (s *objectStorageServiceImpl) WriteJSON(ctx context.Context, bucketName string, objectName string, data interface{}) error {

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(dataBytes)

	// Upload the byte array with PutObject
	n, err := s.client.PutObject(ctx, bucketName, objectName, reader, int64(len(dataBytes)), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}
	fmt.Println("Successfully uploaded bytes: ", n)

	return nil
}
