package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogs()},
		Level: zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey: "level",
			TimeKey: "time",
			MessageKey: "message",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()
}

func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv("LOG_OUTPUT")))

	if output == "" {
		return "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("LOG_LEVEL"))) {
		case "info":
			return zapcore.InfoLevel
		case "error":
			return zapcore.ErrorLevel
		case "debug":
			return zapcore.DebugLevel
		default:
			return zapcore.InfoLevel
	}
}