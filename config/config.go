// Package config deals with server config
package config

import (
	"fmt"
	"sync"
)

// AppConfig defines the configuration structure for the application.
type AppConfig struct {
	Server struct {
		Port string `mapstructure:"port" validate:"required"`
		Env  string `mapstructure:"env" validate:"required"`
	} `mapstructure:"server"`
	CORS struct {
		AllowOrigins []string `mapstructure:"allow_origins"`
	} `mapstructure:"cors"`
	Log struct {
		Level      string `mapstructure:"level"`
		FilePath   string `mapstructure:"file_path"`
		MaxSize    int    `mapstructure:"max_size"`
		MaxBackups int    `mapstructure:"max_backups"`
		MaxAge     int    `mapstructure:"max_age"`
		Compress   bool   `mapstructure:"compress"`
	} `mapstructure:"log"`
}

// Global configuration instance with thread-safe access.
var (
	instance *AppConfig
	once     sync.Once
	loadErr  error
)

// Get returns the global configuration instance.
// This function is safe to call from multiple goroutines.
func Get() (*AppConfig, error) {
	once.Do(func() {
		instance, loadErr = loadConfiguration()
	})
	return instance, loadErr
}

// MustGet returns configuration or panics on error.
func MustGet() *AppConfig {
	cfg, err := Get()
	if err != nil {
		panic(fmt.Sprintf("failed to load configuration: %v", err))
	}
	return cfg
}
