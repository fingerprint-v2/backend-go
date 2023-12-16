package repositories

import (
	"context"

	"github.com/minio/minio-go/v7"
)

type MinioRepository interface {
	CreateBucket(context.Context, string, minio.MakeBucketOptions) error
	UploadObject(context.Context, string, string, string, minio.PutObjectOptions) (*minio.UploadInfo, error)
	DownloadObject(context.Context, string, string, string, minio.GetObjectOptions) error
}

type minioRepositoryImpl struct {
	client *minio.Client
}

func NewMinioRepository(client *minio.Client) MinioRepository {
	return &minioRepositoryImpl{
		client: client,
	}
}

func (r *minioRepositoryImpl) CreateBucket(ctx context.Context, bucketName string, opts minio.MakeBucketOptions) error {
	err := r.client.MakeBucket(ctx, bucketName, opts)
	if err != nil {
		return err
	}
	return nil
}

func (r *minioRepositoryImpl) UploadObject(ctx context.Context, bucketName string, objectName string, path string, opts minio.PutObjectOptions) (*minio.UploadInfo, error) {
	uploadInfo, err := r.client.FPutObject(ctx, bucketName, objectName, path, opts)
	if err != nil {
		return nil, err
	}
	return &uploadInfo, nil
}

func (r *minioRepositoryImpl) DownloadObject(ctx context.Context, bucketName string, objectName string, path string, opts minio.GetObjectOptions) error {
	if err := r.client.FGetObject(ctx, bucketName, objectName, path, opts); err != nil {
		return err
	}
	return nil
}
