package logger

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/pkg/setting"
	"go.uber.org/zap/zapcore"
)

func TestParseLevel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected zapcore.Level
	}{
		{name: "empty defaults to info", input: "", expected: zapcore.InfoLevel},
		{name: "debug", input: "debug", expected: zapcore.DebugLevel},
		{name: "uppercase info", input: "INFO", expected: zapcore.InfoLevel},
		{name: "warning alias", input: "warning", expected: zapcore.WarnLevel},
		{name: "trim spaces", input: " error ", expected: zapcore.ErrorLevel},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			level, err := parseLevel(tt.input)
			if err != nil {
				t.Fatalf("parseLevel() error = %v", err)
			}

			if level != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, level)
			}
		})
	}
}

func TestParseLevelInvalid(t *testing.T) {
	if _, err := parseLevel("invalid"); err == nil {
		t.Fatal("expected error for invalid log level")
	}
}

func TestNewLoggerWithStdout(t *testing.T) {
	zapLogger, err := NewLogger(setting.LoggerSetting{
		Level:  "debug",
		Format: "console",
		Output: "stdout",
	})
	if err != nil {
		t.Fatalf("NewLogger() error = %v", err)
	}

	if zapLogger == nil {
		t.Fatal("expected logger is not nil")
	}

	zapLogger.Info("logger stdout test")
}

func TestNewLoggerWithJSONStdout(t *testing.T) {
	zapLogger, err := NewLogger(setting.LoggerSetting{
		Level:  "debug",
		Format: "json",
		Output: "stdout",
	})
	if err != nil {
		t.Fatalf("NewLogger() error = %v", err)
	}

	zapLogger.Info("logger json stdout test")
}

func TestBuildWriterWithStdout(t *testing.T) {
	writer, err := buildWriter(setting.LoggerSetting{Output: "stdout"})
	if err != nil {
		t.Fatalf("buildWriter() error = %v", err)
	}

	if writer == nil {
		t.Fatal("expected writer is not nil")
	}
}

func TestBuildWriterWithInvalidOutput(t *testing.T) {
	writer, err := buildWriter(setting.LoggerSetting{Output: "invalid"})
	if err == nil {
		t.Fatal("expected error for invalid log output")
	}

	if writer != nil {
		t.Fatal("expected writer is nil when error occurs")
	}
}

func TestBuildWriterWithFile(t *testing.T) {
	logPath := filepath.Join("logs", "logger_test.log")
	_ = os.Remove(logPath)

	writer, err := buildWriter(setting.LoggerSetting{
		Output:   "file",
		FilePath: logPath,
	})
	if err != nil {
		t.Fatalf("buildWriter() error = %v", err)
	}

	if writer == nil {
		t.Fatal("expected writer is not nil")
	}

	if _, err := os.Stat(logPath); err != nil {
		t.Fatalf("expected log file to be created: %v", err)
	}
}
