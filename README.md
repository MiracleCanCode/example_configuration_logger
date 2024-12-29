# my configuration zap logger

## Example usage

### With default config

```go
package main

import (
	"go.uber.org/zap"
)

func main() {
	// Use the default logger configuration
	logger := Logger(DefaultLoggerConfig())
	if logger == nil {
		panic("failed to initialize logger")
	}
	defer logger.Sync() // Ensure proper flushing of logs

	// Example log messages
	logger.Info("Application started with default config")
	logger.Debug("This debug message will not be shown because the default level is Info")
	logger.Warn("A warning using default config")
}
```

### With custom config

```go
package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	customLevel := zapcore.DebugLevel
	loggerConfig := &LoggerConfig{
		Level:            &customLevel,
		Development:      true,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger := Logger(loggerConfig)
	if logger == nil {
		panic("failed to initialize logger")
	}
	defer logger.Sync()

	logger.Info("Application started",
		zap.String("app", "myApp"),
		zap.Int("version", 1),
	)
	logger.Debug("This is a debug message")
	logger.Warn("This is a warning")
	logger.Error("This is an error",
		zap.String("reason", "example reason"),
	)
}
```
