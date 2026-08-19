package setting

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tmpDir := t.TempDir()
	envPath := filepath.Join(tmpDir, ".env")

	content := []byte(`
APP_ENV=development
APP_NAME=go-ecommerce-backend-api
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go_ecommerce
DB_SSLMODE=disable
DB_TIMEZONE=Asia/Ho_Chi_Minh
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
JWT_ACCESS_SECRET=test-access-secret
JWT_REFRESH_SECRET=test-refresh-secret
JWT_ACCESS_TTL_MINUTES=15
JWT_REFRESH_TTL_HOURS=168
LOG_LEVEL=debug
LOG_FORMAT=console
LOG_OUTPUT=stdout
LOG_FILE_PATH=logs/app.log
`)

	if err := os.WriteFile(envPath, content, 0644); err != nil {
		t.Fatal(err)
	}

	config, err := LoadConfig(envPath)
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
