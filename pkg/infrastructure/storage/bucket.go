package storage

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	appconfig "golang-clean-architecture/pkg/config"
)

func CreateBuckets(client *s3.Client) error {
	buckets := []string{
		appconfig.C.Storage.Buckets.WomanImage,
		appconfig.C.Storage.Buckets.BlogImage,
	}

	for _, bucket := range buckets {
		if err := createBucketIfNotExists(client, bucket); err != nil {
			return err
		}
	}
	return nil
}

func createBucketIfNotExists(client *s3.Client, bucket string) error {
	ctx := context.Background()

	// 既に存在するか確認
	_, err := client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err == nil {
		log.Printf("bucket already exists: %s", bucket)
		return nil
	}

	// 存在しなければ作成
	_, err = client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraintApNortheast1,
		},
	})
	if err != nil {
		return err
	}

	log.Printf("bucket created: %s", bucket)
	return nil
}
