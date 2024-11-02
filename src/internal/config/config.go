package config

import (
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	HTTPSettings
}

type HTTPSettings struct {
	Port string
}

// Initialises a default logger from env vars
func InitLogger() {
	slog.Info("Initialising logger")

	logLevel := getLogLevel()
	var addSource bool
	if logLevel == slog.LevelDebug {
		addSource = true
	}

	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: addSource,
		Level:     logLevel,
	})

	slog.SetDefault(slog.New(jsonHandler))

	slog.Info("logger initialised", "level", logLevel.String())
}

func Load() (*Config, error) {
	cfg := Config{}
	cfg.HTTPSettings.Port = getEnvWithDefault("HTTP_PORT", "8080")

	return &cfg, nil
}

// returns slog.Level from env vars
func getLogLevel() slog.Level {
	level := getEnvWithDefault("LOG_LEVEL", "info")
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// checks if an env var exists, defaults to fallback if it does not
func getEnvWithDefault(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		slog.Debug("failed to load environment variable so using default", "key", key)
		return fallback
	}
	slog.Debug("successfully loaded environment variable", "key", key)
	return value
}
