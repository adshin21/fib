// Package config deals with server config
package config

import (
	"log/slog"
	"os"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
		Env  string `mapstructure:"env"`
	} `mapstructure:"server"`
}

var Cfg *Config

func GetConfig() *Config {
	if Cfg == nil {
		slog.Error("config has not been initialized")
		os.Exit(1)
	}
	return Cfg
}
