// Package server is responsible to server initialization with routers
package server

import (
	"log/slog"
	"os"

	"github.com/adshin21/fib/config"
)

func Init() {
	cfg := config.GetConfig()
	r := NewRouter()
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		slog.Error("error running server", "err", err)
		os.Exit(1)
	}
}
