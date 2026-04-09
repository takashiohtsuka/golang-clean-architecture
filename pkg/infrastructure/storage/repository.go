package storage

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	backendoutputport "golang-clean-architecture/pkg/backend/usecase/outputport"
)

type storageRepository struct {
	client *s3.Client
}

func NewStorageRepository(client *s3.Client) backendoutputport.StorageRepository {
	return &storageRepository{client: client}
}

func (r *storageRepository) Upload(ctx context.Context, bucket string, key string, file io.Reader, contentType string) (string, error) {
	_, err := r.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}

	return key, nil
}
