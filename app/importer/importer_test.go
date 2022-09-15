package importer

import (
	"context"
	"testing"
	"time"
)

func TestGo_wrongDir(t *testing.T) {
	if err := Go(context.Background(), "mongodb://localhost:27017", "quake3", "wrong"); err == nil {
		t.Error("expected error, got nil")
	}
}

func TestGo_wrongDbConn(t *testing.T) {
	if err := Go(context.Background(), "", "quake3", "../../data"); err == nil {
		t.Error("expected error, got nil")
	}
}

func TestGo_idleConnection(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := Go(ctx, "mongodb://localhost:27018", "quake3", "../../data"); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestGo_success(t *testing.T) {
	if err := Go(context.Background(), "mongodb://localhost:27017", "quake3", "../../data"); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
