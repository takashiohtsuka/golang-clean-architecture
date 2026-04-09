package outputport

import (
	"context"
	"io"
)

type StorageRepository interface {
	Upload(ctx context.Context, bucket string, key string, file io.Reader, contentType string) (string, error)
}
