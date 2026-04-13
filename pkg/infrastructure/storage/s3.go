package storage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	appconfig "golang-clean-architecture/pkg/config"
)

func NewS3Client() (*s3.Client, error) {
	opts := []func(*config.LoadOptions) error{
		config.WithRegion("ap-northeast-1"),
	}

	// accessKeyIDが設定されている場合はMinIO用の静的認証情報を使用
	// 空の場合はTask Roleの認証情報を自動使用
	if appconfig.C.Storage.AccessKeyID != "" {
		opts = append(opts, config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				appconfig.C.Storage.AccessKeyID,
				appconfig.C.Storage.SecretAccessKey,
				"",
			),
		))
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	var clientOpts []func(*s3.Options)

	// endpointが設定されている場合はMinIO用（ローカル）
	if appconfig.C.Storage.Endpoint != "" {
		clientOpts = append(clientOpts, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(appconfig.C.Storage.Endpoint)
			o.UsePathStyle = true // MinIOはパス形式が必要
		})
	}

	client := s3.NewFromConfig(cfg, clientOpts...)
	return client, nil
}
