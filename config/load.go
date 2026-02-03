package config

import (
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		return nil, err
	}
	return Cfg, nil
}

func Init() {
	if _, err := LoadConfig("./config"); err != nil {
		slog.Error("error loading config", "err", err)
		os.Exit(1)
	}
}
