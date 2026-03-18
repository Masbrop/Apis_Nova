package status

import (
	"context"
	"errors"
	"testing"
)

type repositoryStub struct {
	name string
	err  error
}

func (r repositoryStub) Name() string {
	return r.name
}

func (r repositoryStub) Ping(context.Context) error {
	return r.err
}

func TestServiceCheckReturnsOkWhenDependencyIsAvailable(t *testing.T) {
	service := NewService("apis_nova", "test", repositoryStub{name: "postgres"})

	snapshot, err := service.Check(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if snapshot.Status != "ok" {
		t.Fatalf("expected status ok, got %s", snapshot.Status)
	}

	if snapshot.Dependencies["postgres"] != "up" {
		t.Fatalf("expected postgres dependency to be up, got %s", snapshot.Dependencies["postgres"])
	}
}

func TestServiceCheckReturnsDegradedWhenDependencyFails(t *testing.T) {
	service := NewService("apis_nova", "test", repositoryStub{
		name: "postgres",
		err:  errors.New("dial error"),
	})

	snapshot, err := service.Check(context.Background())
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if snapshot.Status != "degraded" {
		t.Fatalf("expected status degraded, got %s", snapshot.Status)
	}

	if snapshot.Dependencies["postgres"] != "down" {
		t.Fatalf("expected postgres dependency to be down, got %s", snapshot.Dependencies["postgres"])
	}
}
