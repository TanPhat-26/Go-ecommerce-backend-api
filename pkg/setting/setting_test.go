package setting

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("testdata/config.env")
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

	if config.App.Name != "go-ecommerce-backend-api" {
		t.Fatalf("expected app name, got %s", config.App.Name)
	}

	if config.Redis.Addr() != "localhost:6379" {
		t.Fatalf("expected redis addr localhost:6379, got %s", config.Redis.Addr())
	}
}
