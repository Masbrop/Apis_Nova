package status

import (
	"context"
	"fmt"
	"time"
)

type Service struct {
	serviceName string
	environment string
	repository  Repository
}

func NewService(serviceName, environment string, repository Repository) *Service {
	return &Service{
		serviceName: serviceName,
		environment: environment,
		repository:  repository,
	}
}

func (s *Service) Check(ctx context.Context) (Snapshot, error) {
	snapshot := Snapshot{
		Status:      "ok",
		Service:     s.serviceName,
		Environment: s.environment,
		Timestamp:   time.Now().UTC(),
		Dependencies: map[string]string{
			s.repository.Name(): "up",
		},
	}

	if err := s.repository.Ping(ctx); err != nil {
		snapshot.Status = "degraded"
		snapshot.Dependencies[s.repository.Name()] = "down"

		return snapshot, fmt.Errorf("dependency %s is unavailable: %w", s.repository.Name(), err)
	}

	return snapshot, nil
}
