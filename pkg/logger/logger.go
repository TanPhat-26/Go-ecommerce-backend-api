package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/TanPhat-26/Go-ecommerce-backend-api/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(config setting.LoggerSetting) (*zap.Logger, error) {
	level, err := parseLevel(config.Level)
	if err != nil {
		return nil, err
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "message"
	encoderConfig.CallerKey = "caller"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	encoder := buildEncoder(config.Format, encoderConfig)
	writer, err := buildWriter(config)
	if err != nil {
		return nil, err
	}

	core := zapcore.NewCore(encoder, writer, level)

	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	return logger, nil
}

func parseLevel(level string) (zapcore.Level, error) {
	level = strings.ToLower(strings.TrimSpace(level))
	if level == "" {
		level = "info"
	}
	if level == "warning" {
		level = "warn"
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		return zapcore.InfoLevel, fmt.Errorf("unsupported log level: %s", level)
	}
	return zapLevel, nil
}

func buildEncoder(format string, encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
	switch strings.ToLower(format) {
	case "json":
		return zapcore.NewJSONEncoder(encoderConfig)
	case "console", "":
		return zapcore.NewConsoleEncoder(encoderConfig)
	default:
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
}

func buildWriter(config setting.LoggerSetting) (zapcore.WriteSyncer, error) {
	switch strings.ToLower(strings.TrimSpace(config.Output)) {
	case "stdout", "":
		return zapcore.AddSync(os.Stdout), nil

	case "file":
		filePath := config.FilePath
		if filePath == "" {
			filePath = "logs/log.txt"
		}

		if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
			return nil, err
		}

		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, err
		}

		fileWriter := zapcore.AddSync(file)
		consoleWriter := zapcore.AddSync(os.Stdout)
		return zapcore.NewMultiWriteSyncer(consoleWriter, fileWriter), nil
	default:
		return nil, fmt.Errorf("unsupported log output: %s", config.Output)
	}
}
