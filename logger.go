package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerConfig struct {
	Level            *zapcore.Level
	Development      bool
	Encoding         string
	OutputPaths      []string
	ErrorOutputPaths []string
}

func DefaultLoggerConfig() *LoggerConfig {
	defaultLevel := zapcore.InfoLevel
	return &LoggerConfig{
		Level:            &defaultLevel,
		Development:      false,
		Encoding:         "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func Logger(config *LoggerConfig) *zap.Logger {
	if config == nil {
		config = DefaultLoggerConfig()
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(*config.Level),
		Development:      config.Development,
		Encoding:         config.Encoding,
		EncoderConfig:    encoderConfig,
		OutputPaths:      config.OutputPaths,
		ErrorOutputPaths: config.ErrorOutputPaths,
	}

	logger, err := zapConfig.Build()
	if err != nil {
		zap.L().Error("failed to build logger", zap.Error(err))
		return nil
	}

	return logger
}
