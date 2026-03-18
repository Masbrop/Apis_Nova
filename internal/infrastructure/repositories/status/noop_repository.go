package status

import "context"

type NoopRepository struct{}

func NewNoopRepository() NoopRepository {
	return NoopRepository{}
}

func (NoopRepository) Name() string {
	return "runtime"
}

func (NoopRepository) Ping(context.Context) error {
	return nil
}
