package status

import "context"

type Repository interface {
	Name() string
	Ping(ctx context.Context) error
}
