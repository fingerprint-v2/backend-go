package services

import (
	"context"

	"github.com/fingerprint/repositories"
	"github.com/minio/minio-go/v7"
)

type MinioService interface {
	CreateBucket(context.Context, string, minio.MakeBucketOptions) error
	UploadObject(context.Context, string, string, string, minio.PutObjectOptions) (*minio.UploadInfo, error)
	DownloadObject(context.Context, string, string, string, minio.GetObjectOptions) error
}

type minioServiceImpl struct {
	minioRepo repositories.MinioRepository
}

func NewMinioService(minioRepo repositories.MinioRepository) MinioService {
	return &minioServiceImpl{
		minioRepo: minioRepo,
	}
}

func (s *minioServiceImpl) CreateBucket(ctx context.Context, bucketName string, opts minio.MakeBucketOptions) error {
	if err := s.minioRepo.CreateBucket(ctx, bucketName, opts); err != nil {
		return err
	}
	return nil
}

func (s *minioServiceImpl) UploadObject(ctx context.Context, bucketName string, objectName string, path string, opts minio.PutObjectOptions) (*minio.UploadInfo, error) {
	uploadInfo, err := s.minioRepo.UploadObject(ctx, bucketName, objectName, path, opts)
	if err != nil {
		return nil, err
	}
	return uploadInfo, nil
}
func (s *minioServiceImpl) DownloadObject(ctx context.Context, bucketName string, objectName string, path string, opts minio.GetObjectOptions) error {
	if err := s.minioRepo.DownloadObject(ctx, bucketName, objectName, path, opts); err != nil {
		return err
	}
	return nil
}
