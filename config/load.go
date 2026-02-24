package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// ConfigLoader handles the loading of configuration from files and environment variables.
type ConfigLoader struct {
	v *viper.Viper
}

// NewConfigLoader creates a new instance of ConfigLoader.
func NewConfigLoader() *ConfigLoader {
	v := viper.New()
	return &ConfigLoader{v: v}
}

// setDefaults configures all default configuration values.
// All defaults are centralized here for easy maintenance.
func (l *ConfigLoader) setDefaults() {
	// Server defaults
	l.v.SetDefault("server.port", "8080")
	l.v.SetDefault("server.env", "development")

	// CORS defaults
	l.v.SetDefault("cors.allow_origins", []string{"http://localhost:9090"})

	// Log defaults
	l.v.SetDefault("log.level", "info")
	l.v.SetDefault("log.file_path", "app.log")
	l.v.SetDefault("log.max_size", 5)
	l.v.SetDefault("log.max_backups", 10)
	l.v.SetDefault("log.max_age", 14)
	l.v.SetDefault("log.compress", true)
}

// Load reads and unmarshals the configuration.
func (l *ConfigLoader) Load() (*AppConfig, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config" // Default to config directory or file
	}

	// Set all defaults in one place
	l.setDefaults()

	l.v.AddConfigPath(configPath)
	l.v.AddConfigPath("./config") // Fallback
	l.v.SetConfigName("config")
	l.v.SetConfigType("yaml")
	l.v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	l.v.AutomaticEnv()

	if err := l.v.ReadInConfig(); err != nil {
		var configFileNotFound viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFound) {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	var config AppConfig
	if err := l.v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

// loadConfiguration performs the actual configuration loading and validation.
func loadConfiguration() (*AppConfig, error) {
	loader := NewConfigLoader()

	config, err := loader.Load()
	if err != nil {
		return nil, err
	}

	// Validate configuration
	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return config, nil
}
