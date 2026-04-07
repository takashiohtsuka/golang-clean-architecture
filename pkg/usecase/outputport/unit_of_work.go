package outputport

import "context"

type UnitOfWork interface {
	Do(ctx context.Context, fn func() error) error
}
